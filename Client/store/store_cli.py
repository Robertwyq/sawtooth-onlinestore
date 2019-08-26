# coding=utf-8
# Author：Robert WYQ
#
# Date：2019-7-22 ——  2019-8-31
#
# ------------------------------------------------------------------------------
'''     
Command line interface for the simplewallet transaction family.

Parses command line arguments and passes it to the SimpleWalletClient class to process.
'''

import argparse  # python自带的命令行参数解析包
import getpass
import logging   # logging 模块
import os
import sys
import traceback
import pkg_resources

from colorlog import ColoredFormatter

from store.store_client import OnlinestoreClient

DISTRIBUTION_NAME = 'onlinestore'

DEFAULT_URL = 'http://rest-api:8008'

# similar to cookiejar
def create_console_handler(verbose_level):
    '''Setup console logging'''
    del verbose_level # unused
    clog = logging.StreamHandler()
    formatter = ColoredFormatter(
        "%(log_color)s[%(asctime)s %(levelname)-8s%(module)s]%(reset)s "
        "%(white)s%(message)s",
        datefmt="%H:%M:%S",
        reset=True,
        log_colors={
            'DEBUG': 'cyan',
            'INFO': 'green',
            'WARNING': 'yellow',
            'ERROR': 'red',
            'CRITICAL': 'red',
        })

    clog.setFormatter(formatter)
    clog.setLevel(logging.DEBUG)
    return clog

def setup_loggers(verbose_level):
    '''Setup logging.'''
    logger = logging.getLogger()
    logger.setLevel(logging.DEBUG)
    logger.addHandler(create_console_handler(verbose_level))

def add_buy_parser(subparsers, parent_parser):
    '''Define the "buy" command line parsing.'''
    parser = subparsers.add_parser(
        'buy',
        help='buy a certain goods',
        parents=[parent_parser])

    parser.add_argument(
        'amount',
        type=int,
        help='the amount to buy')


    parser.add_argument(
        'address',
        type=str,
        help='the address of store to buy the goods')

    # parser.add_argument(
    #     'goodsName',
    #     type=str,
    #     help='the name of goods')

def add_sell_parser(subparsers, parent_parser):
    '''Define the "sell" command line parsing.'''
    parser = subparsers.add_parser(
        'sell',
        help='sell a certain amount goods in this address',
        parents=[parent_parser])

    parser.add_argument(
        'amount',
        type=int,
        help='the amount to sell')

    parser.add_argument(
        'address',
        type=str,
        help='the address of store to sell the goods')

    # parser.add_argument(
    #     'goodsName',
    #     type=str,
    #     help='the name of goods')

def add_show_parser(subparsers, parent_parser):
    '''Define the "show" command line parsing.'''
    parser = subparsers.add_parser(
        'show',
        help='shows amount in this address',
        parents=[parent_parser])

    parser.add_argument(
        'address',
        type=str,
        help='the address of store')

    # parser.add_argument(
    #     'goodsName',
    #     type=str,
    #     help='the name of goods')

def add_transport_parser(subparsers, parent_parser):
    '''Define the "transport" command line parsing.'''
    parser = subparsers.add_parser(
        'transport',
        help='transport specific goods from one address to the other',
        parents=[parent_parser])

    parser.add_argument(
        'amount',
        type=int,
        help='the amount to transport')

    parser.add_argument(
        'AddressFrom',
        type=str,
        help='the name of address to transport from')
    # 数量减少的一方

    parser.add_argument(
        'AddressTo',
        type=str,
        help='the name of address to transport to')
    # 数量增多的一方

def add_clear_parser(subparsers, parent_parser):

    parser = subparsers.add_parser(
        'empty',
        help='empty the storage',
        parents=[parent_parser])

    parser.add_argument(
        'address',
        type=str,
        help='the address of store')


def create_parent_parser(prog_name):
    '''Define the -V/--version command line options.'''
    # 生成一个parser的类对象，prog是程序名字
    parent_parser = argparse.ArgumentParser(prog=prog_name, add_help=False)

    # 这里有点问题？
    try:
        version = pkg_resources.get_distribution(DISTRIBUTION_NAME).version
    except pkg_resources.DistributionNotFound:
        version = 'UNKNOWN'

    parent_parser.add_argument(
        '-V', '--version',
        action='version',
        version=(DISTRIBUTION_NAME + ' (Hyperledger Sawtooth) version {}')
            .format(version),
        help='display version information')

    return parent_parser


def create_parser(prog_name):
    '''Define the command line parsing for all the options and subcommands.'''
    parent_parser = create_parent_parser(prog_name)
    # description 是 help 前所显示的信息; parents 是 parser对象组成的列表
    parser = argparse.ArgumentParser(
        description='Provides subcommands to manage your simple wallet',
        parents=[parent_parser])
    # dest 是 参数名
    subparsers = parser.add_subparsers(title='subcommands', dest='command')

    subparsers.required = True

    add_buy_parser(subparsers, parent_parser)
    add_sell_parser(subparsers, parent_parser)
    add_show_parser(subparsers, parent_parser)
    add_transport_parser(subparsers, parent_parser)
    add_clear_parser(subparsers, parent_parser)

    return parser

# 这两个get本质上就是去这个地址寻找keyfile的
def _get_keyfile(address):
    '''Get the private key for a customer.'''
    home = os.path.expanduser("~")
    key_dir = os.path.join(home, ".sawtooth", "keys")

    return '{}/{}.priv'.format(key_dir, address)

def _get_pubkeyfile(address):
    '''Get the public key for a customer.'''
    home = os.path.expanduser("~")
    key_dir = os.path.join(home, ".sawtooth", "keys")

    return '{}/{}.pub'.format(key_dir, address)

def do_buy(args):
    '''Implements the "buy" subcommand by calling the client class.'''
    keyfile = _get_keyfile(args.address) # 地址

    client = OnlinestoreClient(baseUrl=DEFAULT_URL, keyFile=keyfile)

    response = client.buy(args.amount)

    print("Response: {}".format(response))

def do_sell(args):
    '''Implements the "sell" subcommand by calling the client class.'''
    keyfile = _get_keyfile(args.address)

    client = OnlinestoreClient(baseUrl=DEFAULT_URL, keyFile=keyfile)

    response = client.sell(args.amount)

    print("Response: {}".format(response))

def do_show(args):
    '''Implements the "show" subcommand by calling the client class.'''
    keyfile = _get_keyfile(args.address)

    client = OnlinestoreClient(baseUrl=DEFAULT_URL, keyFile=keyfile)

    data = client.show()

    if data is not None:
        print("\n{} has a number of = {}\n".format(args.address,
                                                        data.decode()))
    else:
        raise Exception("Data not found: {}".format(args.address))

def do_transport(args):
    '''Implements the "transfer" subcommand by calling the client class.'''
    keyfileFrom = _get_keyfile(args.AddressFrom)
    keyfileTo = _get_pubkeyfile(args.AddressTo)

    clientFrom = OnlinestoreClient(baseUrl=DEFAULT_URL, keyFile=keyfileFrom)

    response = clientFrom.transfer(args.amount, keyfileTo)
    print("Response: {}".format(response))

def do_clear(args):

    keyfile = _get_keyfile(args.address)

    client = OnlinestoreClient(baseUrl=DEFAULT_URL, keyFile=keyfile)

    response = client.clear()
    print("Clear Response: {}".format(response))

def main(prog_name=os.path.basename(sys.argv[0]), args=None):
    '''Entry point function for the client CLI.'''
    if args is None:
        args = sys.argv[1:]
    parser = create_parser(prog_name)
    args = parser.parse_args(args)

    verbose_level = 0

    setup_loggers(verbose_level=verbose_level)

    # Get the commands from cli args and call corresponding handlers
    if args.command == 'buy':
        do_buy(args)
    elif args.command == 'sell':
        do_sell(args)
    elif args.command == 'show':
        do_show(args)
    elif args.command == 'transport':
        # Cannot buy and sell from own account. noop.
        if args.AddressFrom == args.AddressTo:
            raise Exception("Cannot transfer money to self: {}"
                            .format(args.AddressFrom))

        do_transport(args)
    elif args.command == 'empty':
        do_clear(args)
    else:
        raise Exception("Invalid command: {}".format(args.command))

def main_wrapper():
    try:
        main()
    except KeyboardInterrupt:
        pass
    except SystemExit as err:
        raise err
    except BaseException as err:
        traceback.print_exc(file=sys.stderr)
        sys.exit(1)

