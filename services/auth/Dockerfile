FROM julia:1.4-buster

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN julia -e "import Pkg; Pkg.activate(\".\"); Pkg.instantiate()"

CMD julia --project src/main.jl