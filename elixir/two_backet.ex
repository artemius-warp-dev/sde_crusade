defmodule TwoBucket do
  defstruct [:bucket_one, :bucket_two, :moves]
  @type t :: %TwoBucket{bucket_one: integer, bucket_two: integer, moves: integer}

  @doc """
  Find the quickest way to fill a bucket with some amount of water from two buckets of specific sizes.
  """
  @spec measure(
          size_one :: integer,
          size_two :: integer,
          goal :: integer,
          start_bucket :: :one | :two
        ) :: {:ok, TwoBucket.t()} | {:error, :impossible}
  def measure(size_one, size_two, goal, start_bucket) do
    do_measure(
      %TwoBucket{bucket_one: 0, bucket_two: 0, moves: 0},
      {size_one, size_two},
      %{},
      goal,
      start_bucket
    )
    |> List.flatten()
    |> Enum.filter(fn state -> state end)
    |> then(fn
      [] ->
        {:error, :impossible}

      states ->
        {:ok, Enum.min_by(states, fn state -> state.moves end)}
    end)
  end

  defp do_measure(
         %TwoBucket{bucket_one: size_one, bucket_two: 0},
         {size_one, _size_two},
         _visited,
         _goal,
         :two
       ),
       do: false

  defp do_measure(
         %TwoBucket{bucket_one: 0, bucket_two: size_two},
         {_size_one, size_two},
         _visited,
         _goal,
         :one
       ),
       do: false

  defp do_measure(
         %TwoBucket{bucket_one: one, bucket_two: two} = state,
         {size_one, size_two},
         _visited,
         goal,
         _start_bucket
       )
       when (one <= size_one and two <= size_two) and (one == goal or two == goal),
       do: state

  defp do_measure(
         %TwoBucket{bucket_one: one, bucket_two: two, moves: moves} = state,
         {size_one, size_two} = sizes,
         visited,
         goal,
         start_bucket
       )
       when one <= size_one and two <= size_two do

    if visited[{one, two}] do
      false
    else
      u_visited = Map.put(visited, {one, two}, true)

      [
        %{state | bucket_one: 0, bucket_two: two, moves: moves + 1},
        %{state | bucket_one: one, bucket_two: 0, moves: moves + 1},
        %{state | bucket_one: one, bucket_two: size_two, moves: moves + 1},
        %{state | bucket_one: size_one, bucket_two: two, moves: moves + 1},
        %{
          state
          | bucket_one: one + Enum.min([two, size_one - one]),
            bucket_two: two - Enum.min([two, size_one - one]),
            moves: moves + 1
        },
        %{
          state
          | bucket_one: one - Enum.min([one, size_two - two]),
            bucket_two: two + Enum.min([one, size_two - two]),
            moves: moves + 1
        }
      ]
      |> Enum.map(fn state ->
        do_measure(state, sizes, u_visited, goal, start_bucket)
      end)
    end
  end

  defp do_measure(_, _, _, _, _), do: false
end

# do_measure(
#   %{state | bucket_one: 0, bucket_two: two, moves: moves + 1},
#   sizes,
#   u_visited,
#   goal,
#   start_bucket
# ) or
#   do_measure(
#     %{state | bucket_one: one, bucket_two: 0, moves: moves + 1},
#     sizes,
#     u_visited,
#     goal,
#     start_bucket
#   ) or
#   do_measure(
#     %{state | bucket_one: one, bucket_two: size_two, moves: moves + 1},
#     sizes,
#     u_visited,
#     goal,
#     start_bucket
#   ) or
#   do_measure(
#     %{state | bucket_one: size_one, bucket_two: two, moves: moves + 1},
#     sizes,
#     u_visited,
#     goal,
#     start_bucket
#   ) or
#   do_measure(
#     %{
#       state
#       | bucket_one: one + Enum.min([two, 3 - one]),
#         bucket_two: two - Enum.min([two, 3 - one]),
#         moves: moves + 1
#     },
#     sizes,
#     u_visited,
#     goal,
#     start_bucket
#   ) or
#   do_measure(
#     %{
#       state
#       | bucket_one: one - Enum.min([one, 5 - two]),
#         bucket_two: two + Enum.min([one, 5 - two]),
#         moves: moves + 1
#     },
#     sizes,
#     u_visited,
#     goal,
#     start_bucket
#   )

# bucket_one = 3
# bucket_two = 5
# goal = 1
# start_bucket = :two

#FIXME
# bucket_one = 2
# bucket_two = 3
# goal = 3
# start_bucket = :one 

# bucket_one = 5
# bucket_two = 7
# goal = 8
# start_bucket = :one
# expected = {:error, :impossible}

# bucket_one = 6
# bucket_two = 15
# goal = 5
# start_bucket = :one


bucket_one = 3
bucket_two = 5
goal = 1
start_bucket = :one

TwoBucket.measure(bucket_one, bucket_two, goal, start_bucket) |> IO.inspect()
