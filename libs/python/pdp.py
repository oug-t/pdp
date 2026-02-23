from enum import IntEnum

class Level(IntEnum):
    PRIVATE = 0  # No Store, No Train
    PERSONAL = 1 # Store Session, Train User
    GLOBAL = 2   # Store Perm, Train Base

HEADER_NAME = "X-PDP-Level"

def get_header(level: int) -> dict:
    """
    Returns the PDP header dictionary. 
    Defaults to Level.PRIVATE (0) if the input is invalid.
    """
    if level not in [Level.PRIVATE, Level.PERSONAL, Level.GLOBAL]:
        level = Level.PRIVATE
    return {HEADER_NAME: str(level)}
