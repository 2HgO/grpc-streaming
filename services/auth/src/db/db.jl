module db
	import Mongoc

	mutable struct Collections
		User::Union{Mongoc.Collection, Nothing}
	end

	const collections = Collections(nothing)

	function __init__()
		client = Mongoc.Client(ENV["DB_URL"])
		Mongoc.ping(client)

		database = client["auth-server"]

		collections.User = database["users"]

		Mongoc.write_command(database, Mongoc.BSON(
			"createIndexes" => "users",
			"indexes" => [Mongoc.BSON("key" => Mongoc.BSON("username" => 1), "name" => "username_1", "unique" => true)]
		))
	end
end