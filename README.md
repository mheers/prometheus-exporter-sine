# Prometheus Sine Wave Modulator

This project is an example of a Prometheus exporter that is used to modulate a sine wave. The project includes a Go program that modulates the sine wave, a Dockerfile for building a container image for the program, and an example configuration for Prometheus to scrape the exporter every second.

## Usage
1. Build the container image for the Go program using the following commands:

```bash
$ make docker
Start the container using the following commands:
```

```bash
$ docker run -p 8080:8080 mheers/prometheus-exporter-sine
Configure Prometheus to scrape the exporter every second by adding the following configuration to the prometheus.yml file:
```

```yaml
scrape_configs:
  - job_name: 'sine_wave_modulator'
    scrape_interval: 1s
    static_configs:
      - targets: ['localhost:8080']
```

## Developer
This project was developed by [Assistant](https://openai.com/blog/better-language-models/) from OpenAI.


## Transparency
The following commands were given by me:

- `write a go program that runs a prometheus exporter that modulates a sine wave`
- `the program needs to start the prometheus exporter`
- `make the port be configurable`
- `write a Dockerfile to build this`
- `write a scrape configuration that scrapes this every second`
- `write a README for this project that also mentions that you wrote the code and note down all my commands for full transparency`
