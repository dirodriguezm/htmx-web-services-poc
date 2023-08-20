FROM python:3.11-slim as python-base
LABEL org.opencontainers.image.authors="ALeRCE"
ENV PYTHONDONTWRITEBYTECODE=1 \
  PYTHONUNBUFFERED=1 \
  PYTHONFAULTHANDLER=1 \
  PIP_NO_CACHE_DIR=off \
  PIP_DISABLE_PIP_VERSION_CHECK=on \
  PIP_DEFAULT_TIMEOUT=100 \
  POETRY_VIRTUALENVS_IN_PROJECT=true \
  POETRY_NO_INTERACTION=1
RUN pip install poetry

FROM python-base as builder
WORKDIR /app
COPY poetry.lock pyproject.toml ./
RUN poetry install --no-root

FROM python:3.11-slim as app
RUN pip install poetry
COPY --from=builder /app /app
WORKDIR /app
COPY ./src /app/src
COPY ./README.md /app/README.md
COPY ./tailwind.config.js /app/tailwind.config.js
RUN poetry install --only-root
CMD ["poetry", "run", "uvicorn","app.main:create_app", "--host", "0.0.0.0", "--port", "8000"]
