#--------------------------------------#
#             后端编译环境              #
#--------------------------------------#
FROM golang:1.16.4-alpine3.13 AS go_builder

WORKDIR /app/let-diagram-api/
RUN adduser -u 10001 -D app-runner

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on
COPY ./let-diagram-api/go.mod .
COPY ./let-diagram-api/go.sum .
RUN go mod download

COPY ./let-diagram-api .
# 编译源码
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o let-diagram-api .

#--------------------------------------#
#             前端编译环境              #
#--------------------------------------#
FROM node:lts-alpine3.13 AS web_builder
WORKDIR /app/lets-diagram/

# 换源
RUN npm config set registry https://registry.npm.taobao.org
COPY ./lets-diagram .
RUN npm install
RUN npm run build


#--------------------------------------#
#             项目运行环境              #
#--------------------------------------#
FROM nginx:1.19.10-alpine AS server

# COPY --from=go_builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=go_builder /etc/passwd /etc/passwd
# USER app-runner

WORKDIR /app

# 修改时区
RUN apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone

# 复制前后端编译产物
COPY --from=go_builder /app/let-diagram-api/setting.json ./let-diagram-api/setting.json
COPY --from=go_builder /app/let-diagram-api/let-diagram-api ./let-diagram-api/let-diagram-api
COPY --from=web_builder /app/lets-diagram/dist ./lets-diagram/dist
COPY ./lets-diagram/nginx ./lets-diagram/nginx
RUN cp -r ./lets-diagram/nginx/conf/* /etc/nginx 

RUN echo "nginx" >> ./run.sh \
    && echo "./let-diagram-api/let-diagram-api >> ./let-diagram-api/logs/api.log" >> ./run.sh \
    && chmod +x ./run.sh

EXPOSE 8888

CMD ["./run.sh"]


