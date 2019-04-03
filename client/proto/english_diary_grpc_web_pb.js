/**
 * @fileoverview gRPC-Web generated client stub for english_diary
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.english_diary = require('./english_diary_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.english_diary.EnglishDiaryClient =
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

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.english_diary.EnglishDiaryPromiseClient =
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

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.english_diary.GetDiaryRequest,
 *   !proto.english_diary.Diary>}
 */
const methodInfo_EnglishDiary_GetDiary = new grpc.web.AbstractClientBase.MethodInfo(
  proto.english_diary.Diary,
  /** @param {!proto.english_diary.GetDiaryRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.english_diary.Diary.deserializeBinary
);


/**
 * @param {!proto.english_diary.GetDiaryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.english_diary.Diary)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.english_diary.Diary>|undefined}
 *     The XHR Node Readable Stream
 */
proto.english_diary.EnglishDiaryClient.prototype.getDiary =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/english_diary.EnglishDiary/GetDiary',
      request,
      metadata || {},
      methodInfo_EnglishDiary_GetDiary,
      callback);
};


/**
 * @param {!proto.english_diary.GetDiaryRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.english_diary.Diary>}
 *     A native promise that resolves to the response
 */
proto.english_diary.EnglishDiaryPromiseClient.prototype.getDiary =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/english_diary.EnglishDiary/GetDiary',
      request,
      metadata || {},
      methodInfo_EnglishDiary_GetDiary);
};


module.exports = proto.english_diary;

