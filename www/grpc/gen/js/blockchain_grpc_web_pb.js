/**
 * @fileoverview gRPC-Web generated client stub for pactus
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v0.0.0
// source: blockchain.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var transaction_pb = require('./transaction_pb.js')
const proto = {};
proto.pactus = require('./blockchain_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pactus.BlockchainClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pactus.BlockchainPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetBlockRequest,
 *   !proto.pactus.GetBlockResponse>}
 */
const methodDescriptor_Blockchain_GetBlock = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetBlock',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetBlockRequest,
  proto.pactus.GetBlockResponse,
  /**
   * @param {!proto.pactus.GetBlockRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetBlockResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetBlockRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetBlockResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetBlockResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getBlock =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetBlock',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetBlock,
      callback);
};


/**
 * @param {!proto.pactus.GetBlockRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetBlockResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getBlock =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetBlock',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetBlock);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetBlockHashRequest,
 *   !proto.pactus.GetBlockHashResponse>}
 */
const methodDescriptor_Blockchain_GetBlockHash = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetBlockHash',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetBlockHashRequest,
  proto.pactus.GetBlockHashResponse,
  /**
   * @param {!proto.pactus.GetBlockHashRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetBlockHashResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetBlockHashRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetBlockHashResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetBlockHashResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getBlockHash =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetBlockHash',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetBlockHash,
      callback);
};


/**
 * @param {!proto.pactus.GetBlockHashRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetBlockHashResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getBlockHash =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetBlockHash',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetBlockHash);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetBlockHeightRequest,
 *   !proto.pactus.GetBlockHeightResponse>}
 */
const methodDescriptor_Blockchain_GetBlockHeight = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetBlockHeight',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetBlockHeightRequest,
  proto.pactus.GetBlockHeightResponse,
  /**
   * @param {!proto.pactus.GetBlockHeightRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetBlockHeightResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetBlockHeightRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetBlockHeightResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetBlockHeightResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getBlockHeight =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetBlockHeight',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetBlockHeight,
      callback);
};


/**
 * @param {!proto.pactus.GetBlockHeightRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetBlockHeightResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getBlockHeight =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetBlockHeight',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetBlockHeight);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetBlockchainInfoRequest,
 *   !proto.pactus.GetBlockchainInfoResponse>}
 */
const methodDescriptor_Blockchain_GetBlockchainInfo = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetBlockchainInfo',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetBlockchainInfoRequest,
  proto.pactus.GetBlockchainInfoResponse,
  /**
   * @param {!proto.pactus.GetBlockchainInfoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetBlockchainInfoResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetBlockchainInfoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetBlockchainInfoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetBlockchainInfoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getBlockchainInfo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetBlockchainInfo',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetBlockchainInfo,
      callback);
};


/**
 * @param {!proto.pactus.GetBlockchainInfoRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetBlockchainInfoResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getBlockchainInfo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetBlockchainInfo',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetBlockchainInfo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetConsensusInfoRequest,
 *   !proto.pactus.GetConsensusInfoResponse>}
 */
const methodDescriptor_Blockchain_GetConsensusInfo = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetConsensusInfo',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetConsensusInfoRequest,
  proto.pactus.GetConsensusInfoResponse,
  /**
   * @param {!proto.pactus.GetConsensusInfoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetConsensusInfoResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetConsensusInfoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetConsensusInfoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetConsensusInfoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getConsensusInfo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetConsensusInfo',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetConsensusInfo,
      callback);
};


/**
 * @param {!proto.pactus.GetConsensusInfoRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetConsensusInfoResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getConsensusInfo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetConsensusInfo',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetConsensusInfo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetAccountRequest,
 *   !proto.pactus.GetAccountResponse>}
 */
const methodDescriptor_Blockchain_GetAccount = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetAccount',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetAccountRequest,
  proto.pactus.GetAccountResponse,
  /**
   * @param {!proto.pactus.GetAccountRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetAccountResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetAccountRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetAccountResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetAccountResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getAccount =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetAccount',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetAccount,
      callback);
};


/**
 * @param {!proto.pactus.GetAccountRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetAccountResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getAccount =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetAccount',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetAccount);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetValidatorRequest,
 *   !proto.pactus.GetValidatorResponse>}
 */
const methodDescriptor_Blockchain_GetValidator = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetValidator',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetValidatorRequest,
  proto.pactus.GetValidatorResponse,
  /**
   * @param {!proto.pactus.GetValidatorRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetValidatorResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetValidatorRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetValidatorResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetValidatorResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getValidator =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetValidator',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetValidator,
      callback);
};


/**
 * @param {!proto.pactus.GetValidatorRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetValidatorResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getValidator =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetValidator',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetValidator);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetValidatorByNumberRequest,
 *   !proto.pactus.GetValidatorResponse>}
 */
const methodDescriptor_Blockchain_GetValidatorByNumber = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetValidatorByNumber',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetValidatorByNumberRequest,
  proto.pactus.GetValidatorResponse,
  /**
   * @param {!proto.pactus.GetValidatorByNumberRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetValidatorResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetValidatorByNumberRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetValidatorResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetValidatorResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getValidatorByNumber =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetValidatorByNumber',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetValidatorByNumber,
      callback);
};


/**
 * @param {!proto.pactus.GetValidatorByNumberRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetValidatorResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getValidatorByNumber =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetValidatorByNumber',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetValidatorByNumber);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetValidatorAddressesRequest,
 *   !proto.pactus.GetValidatorAddressesResponse>}
 */
const methodDescriptor_Blockchain_GetValidatorAddresses = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetValidatorAddresses',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetValidatorAddressesRequest,
  proto.pactus.GetValidatorAddressesResponse,
  /**
   * @param {!proto.pactus.GetValidatorAddressesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetValidatorAddressesResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetValidatorAddressesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetValidatorAddressesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetValidatorAddressesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getValidatorAddresses =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetValidatorAddresses',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetValidatorAddresses,
      callback);
};


/**
 * @param {!proto.pactus.GetValidatorAddressesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetValidatorAddressesResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getValidatorAddresses =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetValidatorAddresses',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetValidatorAddresses);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetPublicKeyRequest,
 *   !proto.pactus.GetPublicKeyResponse>}
 */
const methodDescriptor_Blockchain_GetPublicKey = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetPublicKey',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetPublicKeyRequest,
  proto.pactus.GetPublicKeyResponse,
  /**
   * @param {!proto.pactus.GetPublicKeyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetPublicKeyResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetPublicKeyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetPublicKeyResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetPublicKeyResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getPublicKey =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetPublicKey',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetPublicKey,
      callback);
};


/**
 * @param {!proto.pactus.GetPublicKeyRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetPublicKeyResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getPublicKey =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetPublicKey',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetPublicKey);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pactus.GetTxPoolContentRequest,
 *   !proto.pactus.GetTxPoolContentResponse>}
 */
const methodDescriptor_Blockchain_GetTxPoolContent = new grpc.web.MethodDescriptor(
  '/pactus.Blockchain/GetTxPoolContent',
  grpc.web.MethodType.UNARY,
  proto.pactus.GetTxPoolContentRequest,
  proto.pactus.GetTxPoolContentResponse,
  /**
   * @param {!proto.pactus.GetTxPoolContentRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pactus.GetTxPoolContentResponse.deserializeBinary
);


/**
 * @param {!proto.pactus.GetTxPoolContentRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pactus.GetTxPoolContentResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pactus.GetTxPoolContentResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pactus.BlockchainClient.prototype.getTxPoolContent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pactus.Blockchain/GetTxPoolContent',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetTxPoolContent,
      callback);
};


/**
 * @param {!proto.pactus.GetTxPoolContentRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pactus.GetTxPoolContentResponse>}
 *     Promise that resolves to the response
 */
proto.pactus.BlockchainPromiseClient.prototype.getTxPoolContent =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pactus.Blockchain/GetTxPoolContent',
      request,
      metadata || {},
      methodDescriptor_Blockchain_GetTxPoolContent);
};


module.exports = proto.pactus;

