const std = @import("std");

pub fn build(b: *std.Build) void {
    // Standard optimization options allow the person running `zig build` to select
    // between Debug, ReleaseSafe, ReleaseFast, and ReleaseSmall. Here we do not
    // set a preferred release mode, allowing the user to decide how to optimize.
    const wasm_target = b.resolveTargetQuery(.{
        .cpu_arch = .wasm32,
        .os_tag = .freestanding,
    });

    const wasm = b.addExecutable(.{
        .name = "app",
        .root_source_file = b.path("models.zig"),
        .target = wasm_target,
        .optimize = .ReleaseSmall,
    });

    wasm.entry = .disabled;
    wasm.rdynamic = true;

    // wasm.installHeader(b.path("include/client.h"), "include/client.h");

    b.installArtifact(wasm);
}
