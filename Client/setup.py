'''
Python package setup (used by Dockerfile).
'''

import os
import subprocess

from setuptools import setup, find_packages

data_files = []

setup(
    name='onlinestore-cli',
    version='1.0',
    description='Sawtooth onlinestore Example',
    author='Robert',
    url='https://github.com/askmish/sawtooth-simplewallet',
    packages=find_packages(),
    install_requires=[
        'aiohttp',
        'colorlog',
        'protobuf',
        'sawtooth-sdk',
        'sawtooth-signing',
        'PyYAML',
    ],
    data_files=data_files,
    entry_points={
        'console_scripts': [
            'onlinestore = store.store_cli:main_wrapper',
        ]
    })

