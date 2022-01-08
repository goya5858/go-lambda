# Cache dependencies & build Go file
FROM golang:1.17.6-buster as build

RUN go env -w GOPROXY=direct
ENV GOPATH=
#GOPATHクリア
ADD ./myModule/ ./
RUN go mod download

# handler.goをmainという名称のファイルにコンパイル
RUN go build -o /main handler.go

# Copy artifacts to clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main
COPY  myModule/SampleModel.onnx /SampleModel.onnx
ENTRYPOINT [ "/main" ]

