# EvolvingCanvas: A Generative Art Framework in Go

EvolvingCanvas is a Go-based framework designed for the creation and manipulation of generative art. It provides a flexible and extensible platform for defining artistic algorithms, rendering them to various output formats, and evolving them through genetic algorithms and other optimization techniques. This project aims to empower artists and developers to explore new frontiers in digital art creation through code.

The core of EvolvingCanvas revolves around the concept of "Art Programs." These programs are defined using a custom domain-specific language (DSL) embedded within Go, allowing users to specify geometric primitives, transformations, color palettes, and composition rules. The framework then interprets and executes these programs to generate visual outputs. Unlike traditional image editing software, EvolvingCanvas offers a procedural approach where art is defined by code, enabling complex and dynamic visual patterns. The ability to version control and automate the art creation process makes it suitable for various applications, from creating unique visual assets for games to generating abstract art pieces for digital display.

One of the most powerful aspects of EvolvingCanvas is its integration with evolutionary algorithms. Users can define fitness functions that evaluate the aesthetic qualities of generated art, guiding the evolution of art programs towards desired outcomes. The framework provides tools for managing populations of art programs, applying genetic operators like mutation and crossover, and selecting individuals based on their fitness scores. This opens up possibilities for discovering novel and unexpected visual styles, pushing the boundaries of artistic expression through computational methods. Furthermore, the modular design allows for easy integration of different rendering backends, enabling support for various output formats and rendering techniques, from simple raster images to complex vector graphics.

EvolvingCanvas goes beyond simple image generation by offering a mechanism for animation. Art programs can be defined with time-dependent parameters, allowing for the creation of dynamic and evolving visual sequences. This opens doors for generating abstract animations, interactive art installations, and dynamic visual effects. The framework provides tools for controlling the animation speed, frame rate, and duration, allowing users to fine-tune the resulting animation. Support for exporting animations to various video formats is planned, making it easy to share and distribute generated art. The integration with genetic algorithms also extends to animation, enabling the evolution of dynamic visual patterns over time.

## Key Features

*   **DSL for Art Program Definition:** EvolvingCanvas utilizes a Go-based DSL to define art programs. This DSL allows for specifying geometric primitives (lines, circles, rectangles, polygons), transformations (translation, rotation, scaling), color palettes, and composition rules using a declarative syntax. This ensures readability and maintainability of the art programs.
*   **Modular Rendering Backend:** The rendering backend is designed to be modular, allowing for easy integration of different rendering techniques. Currently, a basic raster-based renderer is implemented, but future versions will support vector graphics rendering and potentially even 3D rendering. The interface allows developers to easily add their own rendering implementations.
*   **Genetic Algorithm Integration:** The framework provides robust support for genetic algorithms. Users can define fitness functions to evaluate the aesthetic quality of generated art. The framework handles population management, genetic operators (mutation, crossover), and selection mechanisms. The fitness function is a key area of customization and can be tailored to specific artistic goals.
*   **Animation Support:** Art programs can be defined with time-dependent parameters, enabling the creation of dynamic visual sequences. Users can control the animation speed, frame rate, and duration. The framework can be extended to support exporting animations to various video formats.
*   **Configurable Program Parameters:** Art programs accept configurable parameters that influence their behavior. These parameters can be modified manually or controlled by genetic algorithms, allowing for fine-grained control over the generated art. The parameter system supports various data types, including integers, floats, and strings.
*   **Parallel Processing:** The framework leverages Go's concurrency features to parallelize the rendering and evolutionary processes, resulting in significant performance improvements. This allows for the generation of high-resolution images and complex animations in a reasonable amount of time.
*   **Command-Line Interface:** A command-line interface (CLI) is provided for interacting with the framework. This CLI allows users to load art programs, generate images, run evolutionary algorithms, and export animations.

## Technology Stack

*   **Go:** The core programming language. Go provides excellent performance, concurrency features, and a simple syntax, making it well-suited for this project.
*   **Image/Draw Package (Go Standard Library):** Used for basic raster image manipulation and rendering. Future versions will explore more advanced rendering libraries.
*   **Math/Rand Package (Go Standard Library):** Used for random number generation in the genetic algorithm and art program generation.
*   **Flag Package (Go Standard Library):** Used for parsing command-line arguments.

## Installation

1.  **Install Go:** Ensure you have Go installed on your system. You can download the latest version from [https://go.dev/dl/](https://go.dev/dl/).
2.  **Clone the Repository:**
    git clone https://github.com/ezozu/EvolvingCanvas.git
3.  **Navigate to the Project Directory:**
    cd EvolvingCanvas
4.  **Build the Project:**
    go build .
    This will create an executable file named `evolvingcanvas` in the project directory.

## Configuration

EvolvingCanvas uses environment variables for configuration. The following environment variables are supported:

*   `EC_RENDER_WIDTH`: Specifies the width of the generated images in pixels. Defaults to 512.
*   `EC_RENDER_HEIGHT`: Specifies the height of the generated images in pixels. Defaults to 512.
*   `EC_POPULATION_SIZE`: Specifies the size of the population used in the genetic algorithm. Defaults to 100.
*   `EC_MUTATION_RATE`: Specifies the mutation rate used in the genetic algorithm. Defaults to 0.05.
*   `EC_CROSSOVER_RATE`: Specifies the crossover rate used in the genetic algorithm. Defaults to 0.7.

You can set these environment variables in your shell or in a `.env` file in the project directory. Example:

EC_RENDER_WIDTH=1024
EC_RENDER_HEIGHT=1024
EC_POPULATION_SIZE=200

## Usage

To generate an image using a pre-defined art program:

./evolvingcanvas generate -program examples/example.ec -output output.png

To run the genetic algorithm to evolve art programs:

./evolvingcanvas evolve -output_dir evolved_art -generations 100

The `-program` flag specifies the path to the art program file. The `-output` flag specifies the path to the output image file. The `-output_dir` flag specifies the directory to save evolved art programs and images. The `-generations` flag specifies the number of generations to run the genetic algorithm for.

The art program files (e.g., `example.ec`) are plain text files containing the code written in the EvolvingCanvas DSL. The DSL syntax will be documented in detail in a separate document (planned for future release).

## Contributing

We welcome contributions to EvolvingCanvas! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise code with appropriate comments.
4.  Test your changes thoroughly.
5.  Submit a pull request.

Please ensure that your code adheres to the Go coding standards.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/ezozu/EvolvingCanvas/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the Go community for their excellent tools and resources.