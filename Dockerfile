FROM deanxv/coze-discord-proxy:latest

RUN mkdir -p /app/coze-discord-proxy/data/config && chmod 777 /app/coze-discord-proxy/data/config
RUN printf '%s' /etc/secrets/BOT_CONFIG | sed 's/\\/"/g' > /app/coze-discord-proxy/data/config/bot_config.json

RUN echo /etc/secrets/BOT_CONFIG
RUN stat /app/coze-discord-proxy/data/config/bot_config.json
RUN cat /app/coze-discord-proxy/data/config/bot_config.json

WORKDIR /app/coze-discord-proxy/data
EXPOSE 7077

ENTRYPOINT ["/coze-discord-proxy"]
