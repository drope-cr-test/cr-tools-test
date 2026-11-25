<?php
// Helpers for the tools test repo.
function greet(string $name): string {
    return "hello " . $name;
}

function broken_canary(): void {
    undefined_function_call_xyz();
}
