LOGS_DIR=$HOME/rssgo/logs
if [[ ! -e $LOGS_DIR ]]; then
    mkdir -p $LOGS_DIR
elif [[ ! -d $LOGS_DIR ]]; then
    echo "$LOGS_DIR already exists but is not a directory" 1>&2
fi
docker run --name rssgo --network host --log-opt max-size=500m --log-opt max-file=3 -v $HOME/mysql/data:/app/logs -d beegedelow/rsshub