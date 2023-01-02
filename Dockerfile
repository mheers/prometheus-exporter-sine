FROM golang:1.19-alpine as builder

# Create the working directory.
RUN mkdir /app
WORKDIR /app

# Copy the source code and build the application.
COPY . .
RUN go build -o sine_wave_modulator

FROM alpine

# Create the working directory.
RUN mkdir /app
WORKDIR /app

# Copy the compiled binary from the builder image.
COPY --from=builder /app/sine_wave_modulator .

# Run the application.
CMD ["./sine_wave_modulator"]
