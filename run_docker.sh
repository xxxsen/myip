#!/bin/bash

docker run -it -d --name myip --restart=always -p 127.0.0.1:5578:5578 xxxsen/myip 