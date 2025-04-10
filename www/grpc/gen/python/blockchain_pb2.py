# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: blockchain.proto
# Protobuf Python Version: 6.30.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    6,
    30,
    2,
    '',
    'blockchain.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import transaction_pb2 as transaction__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10\x62lockchain.proto\x12\x06pactus\x1a\x11transaction.proto\"-\n\x11GetAccountRequest\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\"C\n\x12GetAccountResponse\x12-\n\x07\x61\x63\x63ount\x18\x01 \x01(\x0b\x32\x13.pactus.AccountInfoR\x07\x61\x63\x63ount\"\x1e\n\x1cGetValidatorAddressesRequest\"=\n\x1dGetValidatorAddressesResponse\x12\x1c\n\taddresses\x18\x01 \x03(\tR\taddresses\"/\n\x13GetValidatorRequest\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\"5\n\x1bGetValidatorByNumberRequest\x12\x16\n\x06number\x18\x01 \x01(\x05R\x06number\"K\n\x14GetValidatorResponse\x12\x33\n\tvalidator\x18\x01 \x01(\x0b\x32\x15.pactus.ValidatorInfoR\tvalidator\"/\n\x13GetPublicKeyRequest\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\"5\n\x14GetPublicKeyResponse\x12\x1d\n\npublic_key\x18\x01 \x01(\tR\tpublicKey\"_\n\x0fGetBlockRequest\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\x12\x34\n\tverbosity\x18\x02 \x01(\x0e\x32\x16.pactus.BlockVerbosityR\tverbosity\"\x83\x02\n\x10GetBlockResponse\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\x12\x12\n\x04hash\x18\x02 \x01(\tR\x04hash\x12\x12\n\x04\x64\x61ta\x18\x03 \x01(\tR\x04\x64\x61ta\x12\x1d\n\nblock_time\x18\x04 \x01(\rR\tblockTime\x12/\n\x06header\x18\x05 \x01(\x0b\x32\x17.pactus.BlockHeaderInfoR\x06header\x12\x34\n\tprev_cert\x18\x06 \x01(\x0b\x32\x17.pactus.CertificateInfoR\x08prevCert\x12)\n\x03txs\x18\x07 \x03(\x0b\x32\x17.pactus.TransactionInfoR\x03txs\"-\n\x13GetBlockHashRequest\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\"*\n\x14GetBlockHashResponse\x12\x12\n\x04hash\x18\x01 \x01(\tR\x04hash\"+\n\x15GetBlockHeightRequest\x12\x12\n\x04hash\x18\x01 \x01(\tR\x04hash\"0\n\x16GetBlockHeightResponse\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\"\x1a\n\x18GetBlockchainInfoRequest\"\xc1\x03\n\x19GetBlockchainInfoResponse\x12*\n\x11last_block_height\x18\x01 \x01(\rR\x0flastBlockHeight\x12&\n\x0flast_block_hash\x18\x02 \x01(\tR\rlastBlockHash\x12%\n\x0etotal_accounts\x18\x03 \x01(\x05R\rtotalAccounts\x12)\n\x10total_validators\x18\x04 \x01(\x05R\x0ftotalValidators\x12\x1f\n\x0btotal_power\x18\x05 \x01(\x03R\ntotalPower\x12\'\n\x0f\x63ommittee_power\x18\x06 \x01(\x03R\x0e\x63ommitteePower\x12H\n\x14\x63ommittee_validators\x18\x07 \x03(\x0b\x32\x15.pactus.ValidatorInfoR\x13\x63ommitteeValidators\x12\x1b\n\tis_pruned\x18\x08 \x01(\x08R\x08isPruned\x12%\n\x0epruning_height\x18\t \x01(\rR\rpruningHeight\x12&\n\x0flast_block_time\x18\n \x01(\x03R\rlastBlockTime\"\x19\n\x17GetConsensusInfoRequest\"\x81\x01\n\x18GetConsensusInfoResponse\x12\x30\n\x08proposal\x18\x01 \x01(\x0b\x32\x14.pactus.ProposalInfoR\x08proposal\x12\x33\n\tinstances\x18\x02 \x03(\x0b\x32\x15.pactus.ConsensusInfoR\tinstances\"Q\n\x17GetTxPoolContentRequest\x12\x36\n\x0cpayload_type\x18\x01 \x01(\x0e\x32\x13.pactus.PayloadTypeR\x0bpayloadType\"E\n\x18GetTxPoolContentResponse\x12)\n\x03txs\x18\x01 \x03(\x0b\x32\x17.pactus.TransactionInfoR\x03txs\"\xdc\x02\n\rValidatorInfo\x12\x12\n\x04hash\x18\x01 \x01(\tR\x04hash\x12\x12\n\x04\x64\x61ta\x18\x02 \x01(\tR\x04\x64\x61ta\x12\x1d\n\npublic_key\x18\x03 \x01(\tR\tpublicKey\x12\x16\n\x06number\x18\x04 \x01(\x05R\x06number\x12\x14\n\x05stake\x18\x05 \x01(\x03R\x05stake\x12.\n\x13last_bonding_height\x18\x06 \x01(\rR\x11lastBondingHeight\x12\x32\n\x15last_sortition_height\x18\x07 \x01(\rR\x13lastSortitionHeight\x12)\n\x10unbonding_height\x18\x08 \x01(\rR\x0funbondingHeight\x12\x18\n\x07\x61\x64\x64ress\x18\t \x01(\tR\x07\x61\x64\x64ress\x12-\n\x12\x61vailability_score\x18\n \x01(\x01R\x11\x61vailabilityScore\"\x81\x01\n\x0b\x41\x63\x63ountInfo\x12\x12\n\x04hash\x18\x01 \x01(\tR\x04hash\x12\x12\n\x04\x64\x61ta\x18\x02 \x01(\tR\x04\x64\x61ta\x12\x16\n\x06number\x18\x03 \x01(\x05R\x06number\x12\x18\n\x07\x62\x61lance\x18\x04 \x01(\x03R\x07\x62\x61lance\x12\x18\n\x07\x61\x64\x64ress\x18\x05 \x01(\tR\x07\x61\x64\x64ress\"\xc4\x01\n\x0f\x42lockHeaderInfo\x12\x18\n\x07version\x18\x01 \x01(\x05R\x07version\x12&\n\x0fprev_block_hash\x18\x02 \x01(\tR\rprevBlockHash\x12\x1d\n\nstate_root\x18\x03 \x01(\tR\tstateRoot\x12%\n\x0esortition_seed\x18\x04 \x01(\tR\rsortitionSeed\x12)\n\x10proposer_address\x18\x05 \x01(\tR\x0fproposerAddress\"\x97\x01\n\x0f\x43\x65rtificateInfo\x12\x12\n\x04hash\x18\x01 \x01(\tR\x04hash\x12\x14\n\x05round\x18\x02 \x01(\x05R\x05round\x12\x1e\n\ncommitters\x18\x03 \x03(\x05R\ncommitters\x12\x1c\n\tabsentees\x18\x04 \x03(\x05R\tabsentees\x12\x1c\n\tsignature\x18\x05 \x01(\tR\tsignature\"\xb1\x01\n\x08VoteInfo\x12$\n\x04type\x18\x01 \x01(\x0e\x32\x10.pactus.VoteTypeR\x04type\x12\x14\n\x05voter\x18\x02 \x01(\tR\x05voter\x12\x1d\n\nblock_hash\x18\x03 \x01(\tR\tblockHash\x12\x14\n\x05round\x18\x04 \x01(\x05R\x05round\x12\x19\n\x08\x63p_round\x18\x05 \x01(\x05R\x07\x63pRound\x12\x19\n\x08\x63p_value\x18\x06 \x01(\x05R\x07\x63pValue\"\x97\x01\n\rConsensusInfo\x12\x18\n\x07\x61\x64\x64ress\x18\x01 \x01(\tR\x07\x61\x64\x64ress\x12\x16\n\x06\x61\x63tive\x18\x02 \x01(\x08R\x06\x61\x63tive\x12\x16\n\x06height\x18\x03 \x01(\rR\x06height\x12\x14\n\x05round\x18\x04 \x01(\x05R\x05round\x12&\n\x05votes\x18\x05 \x03(\x0b\x32\x10.pactus.VoteInfoR\x05votes\"y\n\x0cProposalInfo\x12\x16\n\x06height\x18\x01 \x01(\rR\x06height\x12\x14\n\x05round\x18\x02 \x01(\x05R\x05round\x12\x1d\n\nblock_data\x18\x03 \x01(\tR\tblockData\x12\x1c\n\tsignature\x18\x04 \x01(\tR\tsignature*f\n\x0e\x42lockVerbosity\x12\x18\n\x14\x42LOCK_VERBOSITY_DATA\x10\x00\x12\x18\n\x14\x42LOCK_VERBOSITY_INFO\x10\x01\x12 \n\x1c\x42LOCK_VERBOSITY_TRANSACTIONS\x10\x02*\xa6\x01\n\x08VoteType\x12\x19\n\x15VOTE_TYPE_UNSPECIFIED\x10\x00\x12\x15\n\x11VOTE_TYPE_PREPARE\x10\x01\x12\x17\n\x13VOTE_TYPE_PRECOMMIT\x10\x02\x12\x19\n\x15VOTE_TYPE_CP_PRE_VOTE\x10\x03\x12\x1a\n\x16VOTE_TYPE_CP_MAIN_VOTE\x10\x04\x12\x18\n\x14VOTE_TYPE_CP_DECIDED\x10\x05\x32\x8b\x07\n\nBlockchain\x12=\n\x08GetBlock\x12\x17.pactus.GetBlockRequest\x1a\x18.pactus.GetBlockResponse\x12I\n\x0cGetBlockHash\x12\x1b.pactus.GetBlockHashRequest\x1a\x1c.pactus.GetBlockHashResponse\x12O\n\x0eGetBlockHeight\x12\x1d.pactus.GetBlockHeightRequest\x1a\x1e.pactus.GetBlockHeightResponse\x12X\n\x11GetBlockchainInfo\x12 .pactus.GetBlockchainInfoRequest\x1a!.pactus.GetBlockchainInfoResponse\x12U\n\x10GetConsensusInfo\x12\x1f.pactus.GetConsensusInfoRequest\x1a .pactus.GetConsensusInfoResponse\x12\x43\n\nGetAccount\x12\x19.pactus.GetAccountRequest\x1a\x1a.pactus.GetAccountResponse\x12I\n\x0cGetValidator\x12\x1b.pactus.GetValidatorRequest\x1a\x1c.pactus.GetValidatorResponse\x12Y\n\x14GetValidatorByNumber\x12#.pactus.GetValidatorByNumberRequest\x1a\x1c.pactus.GetValidatorResponse\x12\x64\n\x15GetValidatorAddresses\x12$.pactus.GetValidatorAddressesRequest\x1a%.pactus.GetValidatorAddressesResponse\x12I\n\x0cGetPublicKey\x12\x1b.pactus.GetPublicKeyRequest\x1a\x1c.pactus.GetPublicKeyResponse\x12U\n\x10GetTxPoolContent\x12\x1f.pactus.GetTxPoolContentRequest\x1a .pactus.GetTxPoolContentResponseB:\n\x06pactusZ0github.com/pactus-project/pactus/www/grpc/pactusb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'blockchain_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\006pactusZ0github.com/pactus-project/pactus/www/grpc/pactus'
  _globals['_BLOCKVERBOSITY']._serialized_start=3174
  _globals['_BLOCKVERBOSITY']._serialized_end=3276
  _globals['_VOTETYPE']._serialized_start=3279
  _globals['_VOTETYPE']._serialized_end=3445
  _globals['_GETACCOUNTREQUEST']._serialized_start=47
  _globals['_GETACCOUNTREQUEST']._serialized_end=92
  _globals['_GETACCOUNTRESPONSE']._serialized_start=94
  _globals['_GETACCOUNTRESPONSE']._serialized_end=161
  _globals['_GETVALIDATORADDRESSESREQUEST']._serialized_start=163
  _globals['_GETVALIDATORADDRESSESREQUEST']._serialized_end=193
  _globals['_GETVALIDATORADDRESSESRESPONSE']._serialized_start=195
  _globals['_GETVALIDATORADDRESSESRESPONSE']._serialized_end=256
  _globals['_GETVALIDATORREQUEST']._serialized_start=258
  _globals['_GETVALIDATORREQUEST']._serialized_end=305
  _globals['_GETVALIDATORBYNUMBERREQUEST']._serialized_start=307
  _globals['_GETVALIDATORBYNUMBERREQUEST']._serialized_end=360
  _globals['_GETVALIDATORRESPONSE']._serialized_start=362
  _globals['_GETVALIDATORRESPONSE']._serialized_end=437
  _globals['_GETPUBLICKEYREQUEST']._serialized_start=439
  _globals['_GETPUBLICKEYREQUEST']._serialized_end=486
  _globals['_GETPUBLICKEYRESPONSE']._serialized_start=488
  _globals['_GETPUBLICKEYRESPONSE']._serialized_end=541
  _globals['_GETBLOCKREQUEST']._serialized_start=543
  _globals['_GETBLOCKREQUEST']._serialized_end=638
  _globals['_GETBLOCKRESPONSE']._serialized_start=641
  _globals['_GETBLOCKRESPONSE']._serialized_end=900
  _globals['_GETBLOCKHASHREQUEST']._serialized_start=902
  _globals['_GETBLOCKHASHREQUEST']._serialized_end=947
  _globals['_GETBLOCKHASHRESPONSE']._serialized_start=949
  _globals['_GETBLOCKHASHRESPONSE']._serialized_end=991
  _globals['_GETBLOCKHEIGHTREQUEST']._serialized_start=993
  _globals['_GETBLOCKHEIGHTREQUEST']._serialized_end=1036
  _globals['_GETBLOCKHEIGHTRESPONSE']._serialized_start=1038
  _globals['_GETBLOCKHEIGHTRESPONSE']._serialized_end=1086
  _globals['_GETBLOCKCHAININFOREQUEST']._serialized_start=1088
  _globals['_GETBLOCKCHAININFOREQUEST']._serialized_end=1114
  _globals['_GETBLOCKCHAININFORESPONSE']._serialized_start=1117
  _globals['_GETBLOCKCHAININFORESPONSE']._serialized_end=1566
  _globals['_GETCONSENSUSINFOREQUEST']._serialized_start=1568
  _globals['_GETCONSENSUSINFOREQUEST']._serialized_end=1593
  _globals['_GETCONSENSUSINFORESPONSE']._serialized_start=1596
  _globals['_GETCONSENSUSINFORESPONSE']._serialized_end=1725
  _globals['_GETTXPOOLCONTENTREQUEST']._serialized_start=1727
  _globals['_GETTXPOOLCONTENTREQUEST']._serialized_end=1808
  _globals['_GETTXPOOLCONTENTRESPONSE']._serialized_start=1810
  _globals['_GETTXPOOLCONTENTRESPONSE']._serialized_end=1879
  _globals['_VALIDATORINFO']._serialized_start=1882
  _globals['_VALIDATORINFO']._serialized_end=2230
  _globals['_ACCOUNTINFO']._serialized_start=2233
  _globals['_ACCOUNTINFO']._serialized_end=2362
  _globals['_BLOCKHEADERINFO']._serialized_start=2365
  _globals['_BLOCKHEADERINFO']._serialized_end=2561
  _globals['_CERTIFICATEINFO']._serialized_start=2564
  _globals['_CERTIFICATEINFO']._serialized_end=2715
  _globals['_VOTEINFO']._serialized_start=2718
  _globals['_VOTEINFO']._serialized_end=2895
  _globals['_CONSENSUSINFO']._serialized_start=2898
  _globals['_CONSENSUSINFO']._serialized_end=3049
  _globals['_PROPOSALINFO']._serialized_start=3051
  _globals['_PROPOSALINFO']._serialized_end=3172
  _globals['_BLOCKCHAIN']._serialized_start=3448
  _globals['_BLOCKCHAIN']._serialized_end=4355
# @@protoc_insertion_point(module_scope)
