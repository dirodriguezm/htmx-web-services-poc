from fastapi import APIRouter, Request
from fastapi.templating import Jinja2Templates
from pathlib import Path

template_path = Path(__file__).resolve().parent.parent / "templates"
templates = Jinja2Templates(directory=template_path)

router = APIRouter()


@router.get("/")
def index(request: Request):
    return templates.TemplateResponse("main.html", {"request": request})
