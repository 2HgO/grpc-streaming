include("auth.jl")
include("service.jl")

using gRPC
using ProtoBuf
using Sockets
using .auth

methods = MethodDescriptor[
	MethodDescriptor("Login", 1, UserInfo, Res),
	MethodDescriptor("CreateUser", 2, UserInfo, Res),
	MethodDescriptor("ValidateToken", 3, ValidationInfo, ValidationRes),
]

AuthService = ProtoService(
	ServiceDescriptor("auth.Auth", 1, methods),
	AuthServiceImpl
)

@info "[AuthService] Server starting up..."
run(gRPCServer((AuthService,), ip"0.0.0.0", 55056))
