/**
 * @fileoverview gRPC-Web generated client stub for auth
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.auth = require('./auth_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.auth.AuthClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.auth.AuthPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.UserInfo,
 *   !proto.auth.Res>}
 */
const methodDescriptor_Auth_Login = new grpc.web.MethodDescriptor(
  '/auth.Auth/Login',
  grpc.web.MethodType.UNARY,
  proto.auth.UserInfo,
  proto.auth.Res,
  /**
   * @param {!proto.auth.UserInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Res.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.auth.UserInfo,
 *   !proto.auth.Res>}
 */
const methodInfo_Auth_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto.auth.Res,
  /**
   * @param {!proto.auth.UserInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Res.deserializeBinary
);


/**
 * @param {!proto.auth.UserInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.auth.Res)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.Res>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.Auth/Login',
      request,
      metadata || {},
      methodDescriptor_Auth_Login,
      callback);
};


/**
 * @param {!proto.auth.UserInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.Res>}
 *     Promise that resolves to the response
 */
proto.auth.AuthPromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.Auth/Login',
      request,
      metadata || {},
      methodDescriptor_Auth_Login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.UserInfo,
 *   !proto.auth.Res>}
 */
const methodDescriptor_Auth_CreateUser = new grpc.web.MethodDescriptor(
  '/auth.Auth/CreateUser',
  grpc.web.MethodType.UNARY,
  proto.auth.UserInfo,
  proto.auth.Res,
  /**
   * @param {!proto.auth.UserInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Res.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.auth.UserInfo,
 *   !proto.auth.Res>}
 */
const methodInfo_Auth_CreateUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.auth.Res,
  /**
   * @param {!proto.auth.UserInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.Res.deserializeBinary
);


/**
 * @param {!proto.auth.UserInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.auth.Res)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.Res>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthClient.prototype.createUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.Auth/CreateUser',
      request,
      metadata || {},
      methodDescriptor_Auth_CreateUser,
      callback);
};


/**
 * @param {!proto.auth.UserInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.Res>}
 *     Promise that resolves to the response
 */
proto.auth.AuthPromiseClient.prototype.createUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.Auth/CreateUser',
      request,
      metadata || {},
      methodDescriptor_Auth_CreateUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.auth.ValidationInfo,
 *   !proto.auth.ValidationRes>}
 */
const methodDescriptor_Auth_ValidateToken = new grpc.web.MethodDescriptor(
  '/auth.Auth/ValidateToken',
  grpc.web.MethodType.UNARY,
  proto.auth.ValidationInfo,
  proto.auth.ValidationRes,
  /**
   * @param {!proto.auth.ValidationInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.ValidationRes.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.auth.ValidationInfo,
 *   !proto.auth.ValidationRes>}
 */
const methodInfo_Auth_ValidateToken = new grpc.web.AbstractClientBase.MethodInfo(
  proto.auth.ValidationRes,
  /**
   * @param {!proto.auth.ValidationInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.auth.ValidationRes.deserializeBinary
);


/**
 * @param {!proto.auth.ValidationInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.auth.ValidationRes)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.auth.ValidationRes>|undefined}
 *     The XHR Node Readable Stream
 */
proto.auth.AuthClient.prototype.validateToken =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/auth.Auth/ValidateToken',
      request,
      metadata || {},
      methodDescriptor_Auth_ValidateToken,
      callback);
};


/**
 * @param {!proto.auth.ValidationInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.auth.ValidationRes>}
 *     Promise that resolves to the response
 */
proto.auth.AuthPromiseClient.prototype.validateToken =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/auth.Auth/ValidateToken',
      request,
      metadata || {},
      methodDescriptor_Auth_ValidateToken);
};


module.exports = proto.auth;

