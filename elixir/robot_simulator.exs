defmodule RobotSimulator do
  @type robot() :: any()
  @type direction() :: :north | :east | :south | :west
  @type position() :: {integer(), integer()}

  @directions [:north, :east, :south, :west]

  @doc """
  Create a Robot Simulator given an initial direction and position.

  Valid directions are: `:north`, `:east`, `:south`, `:west`
  """
  @spec create(direction, position) :: robot() | {:error, String.t()}
  def create(direction \\ :north, position \\ {0, 0})

  def create(direction, _) when direction not in @directions do
    {:error, "invalid direction"}
  end

  def create(direction, {x, y} = position) when is_integer(x) and is_integer(y) do
    %{
      position: position,
      direction: direction
    }
  end

  def create(_, _), do: {:error, "invalid position"}

  @doc """
  Simulate the robot's movement given a string of instructions.

  Valid instructions are: "R" (turn right), "L", (turn left), and "A" (advance)
  """
  @spec simulate(robot, instructions :: String.t()) :: robot() | {:error, String.t()}
  def simulate(robot, instructions) do
    instructions
    |> String.graphemes()
    |> Enum.reduce_while(robot, &move/2)
  end

  defp move(instruction, robot) do
    case instruction do
      "R" -> {:cont, %{robot | direction: get_next_direction(robot.direction, 1)}}
      "L" -> {:cont, %{robot | direction: get_next_direction(robot.direction, 3)}}
      "A" -> {:cont, %{robot | position: get_next_position(robot)}}
      _ -> {:halt, {:error, "invalid instruction"}}
    end
  end

  defp get_next_position(robot) do
    {x, y} = robot.position

    case robot.direction do
      :north -> {x, y + 1}
      :south -> {x, y - 1}
      :west -> {x - 1, y}
      :east -> {x + 1, y}
    end
  end

  defp get_next_direction(direction, factor) do
    @directions
    |> Enum.find_index(fn x -> x == direction end)
    |> then(fn i -> Enum.at(@directions, rem(i + factor, 4)) end)
  end

  @doc """
  Return the robot's direction.

  Valid directions are: `:north`, `:east`, `:south`, `:west`
  """
  @spec direction(robot) :: direction()
  def direction(robot) do
    robot.direction
  end

  @doc """
  Return the robot's position.
  """
  @spec position(robot) :: position()
  def position(robot) do
    robot.position
  end
end
