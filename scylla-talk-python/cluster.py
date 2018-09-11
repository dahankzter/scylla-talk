import argparse
import os

from cassandra.cluster import Cluster

parser = argparse.ArgumentParser(description='Do some Scylla stuff!')
parser.add_argument('--host', action="store", dest="host", help='the scylla db host', default="localhost")
parser.add_argument('--schema-dir', action="store", dest="schemadir", help='the schema directory', default="schema")

args = parser.parse_args()

cluster = Cluster([args.host])
session = cluster.connect()

rows = session.execute('SELECT id, name, email, about FROM scylla_talk_ks.user')
for user_row in rows:
    print(user_row.id, user_row.name, user_row.email, user_row.about)