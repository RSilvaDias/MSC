import os
import sys
import subprocess

def main():
    if len(sys.argv) < 3:
        print("Usage: python compiler.py <source_file_path> <output_file_path>")
        return

    source_file_path = sys.argv[1]
    output_file_path = sys.argv[2]

    print(f"Using source file '{source_file_path}'.")

    # Check if the source file exists
    if not os.path.exists(source_file_path):
        print(f"Error: Source file '{source_file_path}' does not exist.")
        return

    # Determine the file extension
    _, ext = os.path.splitext(source_file_path)

    # Determine the compilation command based on the file extension
    if ext == ".go":
        # Compile the source file to WebAssembly using TinyGo
        cmd = ["tinygo", "build", "-o", output_file_path, "-target=wasi", source_file_path]
    elif ext == ".cpp":
        # Compile the source file to WebAssembly using Binaryen (via emcc)
        cmd = ["emcc", source_file_path, "-o", output_file_path, "-s", "WASM=1"]
    else:
        print(f"Error: Unsupported file extension '{ext}'. Supported extensions are .go and .cpp.")
        return

    # Run the compilation command
    try:
        subprocess.run(cmd, check=True)
        
        # Check if the output file exists
        if os.path.exists(output_file_path):
            print(f"Compiled '{source_file_path}' to WebAssembly binary '{output_file_path}' successfully.")
        else:
            print(f"Error: Compilation succeeded but output file '{output_file_path}' is missing.")
    except subprocess.CalledProcessError as e:
        print(f"Error: Failed to compile '{source_file_path}' to WebAssembly: {e}")

if __name__ == "__main__":
    main()
