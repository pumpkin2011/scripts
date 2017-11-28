#!/bin/bash
psql postgres << EOF
	drop database "aico-nexus_development";
	create database "aico-nexus_development";
EOF
psql aico-nexus_development < /Users/jude/Downloads/aico-nexus_dump_2017-11-20.pgsql
