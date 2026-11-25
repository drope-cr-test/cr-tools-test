"""Marker Pylint plugin for C-T3: registers a checker that emits a message
containing CR-MARKER-T3 when it loads. If this plugin loads, the marker
appears in CodeRabbit's Pylint findings."""
from pylint.checkers import BaseChecker


class MarkerChecker(BaseChecker):
    name = "cr-marker"
    msgs = {
        "W9901": (
            "CR-MARKER-T3-7f3a9 plugin loaded",
            "cr-marker-loaded",
            "Marker: the load-plugins path executed.",
        ),
    }

    def visit_module(self, node):
        self.add_message("cr-marker-loaded", node=node)


def register(linter):
    linter.register_checker(MarkerChecker(linter))
