defmodule Bowling do
  @doc """
    Creates a new game of bowling that can be used to store the results of
    the game
  """

  defstruct pins: 10,
            balls: 1,
            score: 0,
            frame: 0,
            strike?: false,
            previous_balls: [],
            spare?: false,
            fill_balls: 0

  @spec start() :: any
  def start do
    %Bowling{}
  end

  @doc """
    Records the number of pins knocked down on a single roll. Returns `any`
    unless there is something wrong with the given number of pins, in which
    case it returns a helpful error tuple.
  """

  @spec roll(any, integer) :: {:ok, any} | {:error, String.t()}

  def roll(_, roll) when roll < 0, do: {:error, "Negative roll is invalid"}
  def roll(_, roll) when roll > 10, do: {:error, "Pin count exceeds pins on the lane"}

  def roll(%Bowling{frame: 10, fill_balls: 0}, _) do
    {:error, "Cannot roll after game is over"}
  end

  def roll(%Bowling{frame: 10, fill_balls: count, pins: pins} = game, roll)
      when pins - roll >= 0 do
    previous_balls = game.previous_balls
    strike_score = Enum.sum(previous_balls)

    {balls, _frame} = update_frame(game, roll)

    strike? =
      if 10 in previous_balls do
        true
      else
        false
      end

    previous_balls =
      if strike? do
        [roll | List.delete(previous_balls, 10)]
      else
        []
      end

    {:ok,
     %{
       game
       | score: roll + game.score + strike_score,
         previous_balls: previous_balls,
         strike?: strike?,
         pins: update_pins(game, roll, balls),
         fill_balls: count - 1
     }}
  end

  def roll(%Bowling{pins: pins}, roll)
      when pins - roll < 0,
      do: {:error, "Pin count exceeds pins on the lane"}

  def roll(%Bowling{strike?: false} = game, 10) do
    {balls, frame} = update_frame(game, 10)

    {:ok,
     %Bowling{
       score: game.score + 10,
       strike?: true,
       frame: frame,
       fill_balls: fill_balls?(frame, 10, game),
       balls: balls
     }}
  end

  def roll(%Bowling{strike?: true, previous_balls: previous_balls} = game, roll)
      when length(previous_balls) < 2 do
    {balls, frame} = update_frame(game, roll)


    {:ok,
     %{
       game
       | previous_balls: [roll | game.previous_balls],
         score: roll + game.score,
         balls: balls,
         frame: frame,
         pins: update_pins(game, roll, balls),
         fill_balls: fill_balls?(frame, roll, game)
     }}
  end

  def roll(%Bowling{strike?: true, previous_balls: previous_balls} = game, roll)
      when length(previous_balls) == 2 do
    strike_score = Enum.sum(previous_balls)
    {balls, frame} = update_frame(game, roll)

    strike? =
      if 10 in previous_balls do
        true
      else
        false
      end

    previous_balls =
      if strike? do
        [roll | List.delete(previous_balls, 10)]
      else
        []
      end

    {:ok,
     %{
       game
       | strike?: strike?,
         previous_balls: previous_balls,
         score: roll + game.score + strike_score,
         balls: balls,
         frame: frame,
         pins: update_pins(game, roll, balls),
         fill_balls: fill_balls?(frame, roll, game)
     }}
  end

  def roll(%Bowling{pins: pins, score: score} = game, roll)
      when pins - roll == 0 and roll > 0 do

    {balls, frame} = update_frame(game, roll)

    {spare_points, previous_balls} =
      if length(game.previous_balls) == 1 do
        {game.previous_balls |> hd, []}
      else
        {0, []}
      end

    {:ok,
     %Bowling{
       spare?: true,
       score: score + roll + spare_points,
       previous_balls: previous_balls,
       balls: balls,
       frame: frame,
       pins: update_pins(game, roll, balls),
       fill_balls: fill_balls?(frame, roll, game)
     }}
  end

  def roll(
        %Bowling{spare?: true, previous_balls: previous_balls} = game,
        roll
      )
      when length(previous_balls) == 1 do

    {balls, frame} = update_frame(game, roll)

    pins_remainder = update_pins(game, roll, balls)

    {:ok,
     %{
       game
       | spare?: false,
         previous_balls: [],
         score: game.score + roll + hd(previous_balls),
         pins: pins_remainder,
         balls: balls,
         frame: frame,
         fill_balls: fill_balls?(frame, roll, game)
     }}
  end

  def roll(
        %Bowling{spare?: true, previous_balls: previous_balls} = game,
        roll
      )
      when length(previous_balls) < 1 do

    {balls, frame} = update_frame(game, roll)

    pins_remainder = update_pins(game, roll, balls)

    {:ok,
     %{
       game
       | previous_balls: [roll | game.previous_balls],
         score: game.score + roll,
         balls: balls,
         frame: frame,
         pins: pins_remainder,
         fill_balls: fill_balls?(frame, roll, game)
     }}
  end

  def roll(game, roll) do
    {balls, frame} = update_frame(game, roll)

    pins_remainder = update_pins(game, roll, balls)

    state = %{
      game
      | pins: pins_remainder,
        score: game.score + roll,
        balls: balls,
        frame: frame,
        fill_balls: fill_balls?(frame, roll, game)
    }

    {:ok, state}
  end

  defp update_frame(game, roll) when roll != 10 do
    res =
      case game.balls do
        2 -> {1, update_frame(game)}
        _ -> {game.balls + 1, game.frame}
      end

    res
  end

  defp update_frame(game, 10) do
    res = {1, update_frame(game)}
    res
  end

  defp update_frame(game) do
    case game do
      %Bowling{frame: 10, fill_balls: true} -> game.frame
      _ -> game.frame + 1
    end
  end

  defp update_pins(game, roll, balls) do
    case roll do
      10 -> 10
      roll -> if balls == 1, do: 10, else: game.pins - roll
    end
  end

  defp fill_balls?(10, roll, %Bowling{pins: pins}) when roll == 10 or roll - pins == 0 do
    cond do
      roll == 10 -> 2
      roll - pins == 0 -> 1
    end
  end

  defp fill_balls?(_, _, game), do: game.fill_balls

  @doc """
    Returns the score of a given game of bowling if the game is complete.
    If the game isn't complete, it returns a helpful error tuple.
  """

  @spec score(any) :: {:ok, integer} | {:error, String.t()}
  def score(%Bowling{frame: 10, score: score, fill_balls: 0}) do
    {:ok, score}
  end

  def score(_game) do
    {:error, "Score cannot be taken until the end of the game"}
  end
end

roll_reduce = fn game, rolls ->
  Enum.reduce(rolls, game, fn roll, game ->
    {:ok, game} = Bowling.roll(game, roll)

    game
  end)
end

game = Bowling.start()

# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
# rolls = [3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6]
# rolls = [6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
# rolls = [6, 4, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
# rolls = [5, 5, 3, 7, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3, 7]
# rolls = [10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
# rolls = [10, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
# rolls = [10, 10, 10, 5, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 7, 1]
rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 7, 3]
# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 10]
# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 0, 1]
# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 3, 10]
# rolls = [10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10]
# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 10]
# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 10, 0, 1]

# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 5]

game = roll_reduce.(game, rolls)
Bowling.score(game) |> IO.inspect()

# Bowling.roll(game, -1) == {:error, "Negative roll is invalid"}
# Bowling.roll(game, 11)
# |> IO.inspect()

# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 5]
# game = roll_reduce.(game, rolls)
# Bowling.roll(game, 6) |> IO.inspect()

## {:ok, game} = Bowling.roll(game, 5)

# Bowling.roll(game, 6)
# |> IO.inspect()

# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10]

# game = roll_reduce.(game, rolls)

# Bowling.roll(game, 11) #== {:error, "Pin count exceeds pins on the lane"}
# |> IO.inspect()

# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1]
# game = roll_reduce.(game, rolls)
# Bowling.roll(game, 0) |> IO.inspect()

# rolls = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 5]
# game = roll_reduce.(game, rolls)
# Bowling.roll(game, 6) |> IO.inspect()
