/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = require('./source_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.SourceClient =
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
proto.SourcePromiseClient =
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
 *   !proto.VideoInfo,
 *   !proto.Video>}
 */
const methodDescriptor_Source_GetVideo = new grpc.web.MethodDescriptor(
  '/Source/GetVideo',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.VideoInfo,
  proto.Video,
  /**
   * @param {!proto.VideoInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Video.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.VideoInfo,
 *   !proto.Video>}
 */
const methodInfo_Source_GetVideo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.Video,
  /**
   * @param {!proto.VideoInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.Video.deserializeBinary
);


/**
 * @param {!proto.VideoInfo} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.Video>}
 *     The XHR Node Readable Stream
 */
proto.SourceClient.prototype.getVideo =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/Source/GetVideo',
      request,
      metadata || {},
      methodDescriptor_Source_GetVideo);
};


/**
 * @param {!proto.VideoInfo} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.Video>}
 *     The XHR Node Readable Stream
 */
proto.SourcePromiseClient.prototype.getVideo =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/Source/GetVideo',
      request,
      metadata || {},
      methodDescriptor_Source_GetVideo);
};


module.exports = proto;

