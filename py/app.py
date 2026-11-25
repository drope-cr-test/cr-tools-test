"""Trivial module for the Pylint tests."""


def greet(name):
    return "hello " + name


# C-T round 2 canary: undefined name proves Pylint ran.
def broken():
    return undefined_name_here
