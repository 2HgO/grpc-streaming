module AuthServiceImpl
	using ProtoBuf
	using ..auth: UserInfo, Res, ValidationRes, ValidationInfo
	using Mongoc: BSON, BSONObjectId, BSONError, find_one
	import Bcrypt: GenerateFromPassword, CompareHashAndPassword
	import JSONWebTokens: JSONWebTokens, HS256, encode, decode
	import Dates: now, Day, UTC, DateTime
	using gRPC
	include("db/db.jl"); import .db

	const encoding = HS256("static_secret")

	function Login(message::UserInfo) :: Res
		username::String = has_field(message, :username) ? get_field(message, :username) : throw(gRPCException(InvalidArgument, "username not defined"))
		password::String = has_field(message, :password) ? get_field(message, :password) : throw(gRPCException(InvalidArgument, "password not defined"))
		(length(username) < 4) && throw(gRPCException(InvalidArgument, "username too short"))
		(length(password) < 4) && throw(gRPCException(InvalidArgument, "password too short"))

		res = find_one(db.collections.User, BSON("username" => username))
		isnothing(res) && return Res(success=false, token="")

		!CompareHashAndPassword(res["password"], password) && return Res(success=false, token="")

		claims = Dict("user" => string(res["_id"]), "exp" => string(now(UTC) + Day(1)))
		token = encode(encoding, claims)

		Res(success=true, token=token)
	end

	function CreateUser(message::UserInfo) :: Res
		try
			username::String = has_field(message, :username) ? get_field(message, :username) : throw(gRPCException(InvalidArgument, "username not defined"))
			password::String = has_field(message, :password) ? get_field(message, :password) : throw(gRPCException(InvalidArgument, "password not defined"))
			(length(username) < 4) && throw(gRPCException(InvalidArgument, "username too short"))
			(length(password) < 4) && throw(gRPCException(InvalidArgument, "password too short"))

			hashedPassword = String(GenerateFromPassword(password, 10))

			res = push!(db.collections.User, BSON("createdAt" => now(UTC), "username" => username, "password" => hashedPassword))

			claims = Dict("user" => string(res.inserted_oid), "exp" => string(now(UTC) + Day(1)))
			token = encode(encoding, claims)

			return Res(success=true, token=token)
		catch ex
			isa(ex, BSONError) && occursin("dup key", cstring_to_jstring(ex.message)) && rethrow(gRPCException(AlreadyExists, "username already in use"))
			rethrow(ex)
		end
	end

	function ValidateToken(message::ValidationInfo) :: ValidationRes
		token = has_field(message, :token) ? get_field(message, :token) : throw(gRPCException(InvalidArgument, "token not defined"))

		claims = decode(encoding, token)

		(!haskey(claims, "user") || !haskey(claims, "exp")) && throw(gRPCException(InvalidArgument, "invalid token")) #errors.InvalidTokenError())
		exp = nothing
		try
			exp = DateTime(string(claims["exp"]))
		catch ex
			rethrow(gRPCException(InvalidArgument, "invalid token"))
		end

		(now() > exp) && throw(gRPCException(InvalidArgument, "expired token"))

		id = nothing
		try
			id = BSONObjectId(string(claims["user"]))
		catch ex
			rethrow(gRPCException(InvalidArgument, "invalid token"))
		end
		
		res = find_one(db.collections.User, BSON("_id" => id))

		ValidationRes(success = !isnothing(res))
	end

	function cstring_to_jstring(cstring::NTuple{504, UInt8})
		s = ""
		for c in cstring
			c_char = Char(c)
			if c_char == '\0'
				break
			else
				s = s * c_char
			end
		end
		return s
	end

	gRPC.wrap(::JSONWebTokens.MalformedJWTError) = gRPCException(InvalidArgument, "invalid jwt token")
	gRPC.wrap(::JSONWebTokens.InvalidSignatureError) = gRPCException(InvalidArgument, "invalid jwt token")
	gRPC.wrap(::JSONWebTokens.NotSupportedJWTError) = gRPCException(InvalidArgument, "invalid jwt token")
	gRPC.wrap(b::BSONError) = gRPCException(cstring_to_jstring(b.message))
	gRPC.wrap(e::TypeError) = gRPCException(InvalidArgument, string(e.context))
	gRPC.wrap(e::ErrorException) = gRPCException(string(e.msg))
	gRPC.wrap(::Union{MethodError, ArgumentError, EOFError}) = gRPCException(InvalidArgument, "validation error")
end
