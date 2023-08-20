from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from pathlib import Path
from .routes import router
from fastapi.middleware.cors import CORSMiddleware


src_path = Path(__file__).resolve().parent.parent


def create_app() -> FastAPI:
    """Create a FastAPI app with the specified settings."""
    app = FastAPI()
    origins = [
        "http://localhost:8000",
        "http://localhost:8001",
    ]
    app.add_middleware(
        CORSMiddleware,
        allow_origins=origins,
        allow_methods=["GET"],
        allow_headers=["*"],
    )

    app.mount(
        "/static",
        StaticFiles(directory=src_path / "static"),
        name="static",
    )

    app.include_router(router)

    return app
