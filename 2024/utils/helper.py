import datetime
import pathlib
import sys
from zoneinfo import ZoneInfo

import requests


__all__ = ("InputHelper",)


with open(".TOKEN") as fp:
    TOKEN: str = fp.read()

if not TOKEN:
    raise RuntimeError("No SESSION_TOKEN found.")


BASE_URL: str = "https://adventofcode.com/{year}/day/{day}/input"


class InputHelper:
    def __init__(self, *, year: int | None = None, day: int | None = None) -> None:
        now: datetime.datetime = datetime.datetime.now(tz=ZoneInfo("America/New_York"))
        if not year and not day and now.month != 12:
            raise RuntimeError("You must provide a year and day when running outside of December.")

        self.year: int = year or now.year
        self.day: int = day or now.day

        self.session: requests.Session = requests.Session()
        self.session.headers.update(self.headers)

        self.data: str | None = None

    @property
    def headers(self) -> dict[str, str]:
        pyver = f"{sys.version_info[0]}.{sys.version_info[1]}"

        return {
            "User-Agent": f"https://github.com/EvieePy/AoC (evieepy@gmail.com) Python/{pyver} requests/{requests.__version__}",
            "Cookie": f"session={TOKEN}",
        }

    def fetch(self, renew: bool = False) -> None:
        path = pathlib.Path(f"inputs")

        if not renew and path.exists():
            try:
                with open(path / f"input_{self.day}.txt") as fp:
                    self.data = fp.read()
            except (FileNotFoundError, OSError):
                pass
            else:
                return

        with self.session.get(BASE_URL.format(year=self.year, day=self.day)) as resp:
            resp.raise_for_status()
            self.data = resp.text

        if not self.data:
            raise RuntimeError("No input data found.")

        path.mkdir(parents=True, exist_ok=True)
        with open(path / f"input_{self.day}.txt", "w") as fp:
            fp.write(self.data)

    def as_lines(self) -> list[str]:
        if not self.data:
            return []

        return self.data.splitlines()
