# Cache dependencies & build Go file
FROM public.ecr.aws/lambda/provided:al2 as build

RUN yum install -y golang
RUN go env -w GOPROXY=direct

ADD ./TestLambda/ ./
RUN go mod download
# handler.goをmainという名称のファイルにコンパイル
RUN go build -o /main handler.go

# Copy artifacts to clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]

