import os
from pathlib import Path

HOME_DIR = Path.home() 
COMDEX_DIR_PATH  = os.path.abspath(os.curdir)

NODE_MONIKER = "testdev"
CHAIN_ID = "test-1"
GENESIS_ACCOUNT_NAME = "cooluser"
GENESIS_TOKENS = "1000000000000000000000stake,100000000000000000000000ucmdx,100000000000000000000000ucmst,100000000000000000000000uosmo,100000000000000000000000uatom"

VOTING_PERIOD_IN_SEC = 10
DEPOSIT_PERIOD_IN_SEC = 10