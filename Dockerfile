FROM deanxv/coze-discord-proxy:latest

RUN find / -name "bot_config" -print
RUN mkdir -p /app/coze-discord-proxy/data/config && chmod 777 /app/coze-discord-proxy/data/config
RUN cat /etc/secrets/BOT_CONFIG > /app/coze-discord-proxy/data/config/bot_config.json

RUN cat /etc/secrets/BOT_CONFIG
RUN stat /app/coze-discord-proxy/data/config/bot_config.json
RUN cat /app/coze-discord-proxy/data/config/bot_config.json

WORKDIR /app/coze-discord-proxy/data
EXPOSE 7077

ENTRYPOINT ["/coze-discord-proxy"]
