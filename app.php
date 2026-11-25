<?php
// Trivial file for the PHPStan tests (root-level control).
function greet(string $name): string {
    return "hello " . unknown_helper($name);
}
