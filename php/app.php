<?php
// Trivial file for the PHPStan tests.
function greet(string $name): string {
    return "hello " . $name;
}

// C-T round 2 canary: call to an undefined function (PHPStan level 1).
function broken_canary(): void {
    undefined_function_call_xyz();
}
