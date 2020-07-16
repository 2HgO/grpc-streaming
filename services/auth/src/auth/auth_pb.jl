# syntax: proto3
using ProtoBuf
import ProtoBuf.meta

mutable struct UserInfo <: ProtoType
    username::AbstractString
    password::AbstractString
    UserInfo(; kwargs...) = (o=new(); fillunset(o); isempty(kwargs) || ProtoBuf._protobuild(o, kwargs); o)
end #mutable struct UserInfo

mutable struct Res <: ProtoType
    success::Bool
    token::AbstractString
    Res(; kwargs...) = (o=new(); fillunset(o); isempty(kwargs) || ProtoBuf._protobuild(o, kwargs); o)
end #mutable struct Res

mutable struct ValidationInfo <: ProtoType
    token::AbstractString
    ValidationInfo(; kwargs...) = (o=new(); fillunset(o); isempty(kwargs) || ProtoBuf._protobuild(o, kwargs); o)
end #mutable struct ValidationInfo

mutable struct ValidationRes <: ProtoType
    success::Bool
    ValidationRes(; kwargs...) = (o=new(); fillunset(o); isempty(kwargs) || ProtoBuf._protobuild(o, kwargs); o)
end #mutable struct ValidationRes

export UserInfo, Res, ValidationInfo, ValidationRes
