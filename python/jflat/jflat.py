#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# echo '{"key1": {"key2": {"key3":true}}}' | ./jflat.py
# ./jflat.py sample.json

import sys
import select
import json

def flatten_json(j):
    out={}
    def flatten(x, name=''):
        # If the nested key-value is of dict type
        if type(x) is dict:
            for a in x:
                flatten(x[a], name + a + '.')

        # If the nested key-value is of list type
        elif type(x) is list:
            i = 0
            for a in x:
                flatten(a, name + str(i) + '.')
                i += 1
        else:
            out[name[:-1]] = x

    flatten(j)
    return out

if __name__ == '__main__':

    if not select.select([sys.stdin,],[],[],0.0)[0]:
        with open(sys.argv[1]) as jsonfile:
            jdata = json.load(jsonfile)
    else:
        jdata = json.load(sys.stdin)

    flat = flatten_json(jdata)
    print(json.dumps(flat, indent=2))

