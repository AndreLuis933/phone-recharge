import logging
import sys
from pathlib import Path

is_docker = Path("/.dockerenv").exists() or Path("/run/.containerenv").exists()

if not is_docker:
    from dotenv import load_dotenv

    load_dotenv()

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s",
    stream=sys.stdout,
    force=True,
)


sys.stdout.reconfigure(line_buffering=True)
sys.stderr.reconfigure(line_buffering=True)
