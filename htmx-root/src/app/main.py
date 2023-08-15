from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from pathlib import Path
from .routes import router

src_path = Path(__file__).resolve().parent.parent


def create_app() -> FastAPI:
    """Create a FastAPI app with the specified settings."""
    app = FastAPI()

    app.mount(
        "/static",
        StaticFiles(directory=src_path / "static"),
        name="static",
    )

    app.include_router(router)

    return app
