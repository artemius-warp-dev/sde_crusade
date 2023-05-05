defmodule ComplexNumbers do
  @typedoc """
  In this module, complex numbers are represented as a tuple-pair containing the real and
  imaginary parts.
  For example, the real number `1` is `{1, 0}`, the imaginary number `i` is `{0, 1}` and
  the complex number `4+3i` is `{4, 3}'.
  """
  @type complex :: {float, float}

  @doc """
  Return the real part of a complex number
  """
  @spec real(a :: complex) :: float
  def real({real, _imaginary}) do
    real
  end

  @doc """
  Return the imaginary part of a complex number
  """
  @spec imaginary(a :: complex) :: float
  def imaginary({_, imaginary}) do
    imaginary
  end

  @doc """
  Multiply two complex numbers, or a real and a complex number
  """
  @spec mul(a :: complex | float, b :: complex | float) :: complex
  def mul({r_a, i_a} = a, {r_b, i_b} = b) do
    {(real(a) * real(b) - imaginary(a) * imaginary(b)),  (imaginary(a) * real(b) + real(a)*imaginary(b))}
  end

  def mul({r_a, i_a}, b), do: {r_a * b, i_a * b} 
  def mul(a, {r_b, i_b}), do: {r_b * a, i_b * a}

  @doc """
  Add two complex numbers, or a real and a complex number
  """
  @spec add(a :: complex | float, b :: complex | float) :: complex
  def add({r_a, i_a}, {r_b, i_b}) do
    {r_b + r_a, i_a + i_b}
  end
  def add({r_a, i_a}, b), do: {r_a + b, i_a}
  def add(a, {r_b, i_b}), do: {r_b + a, i_b}
  

  @doc """
  Subtract two complex numbers, or a real and a complex number
  """
  @spec sub(a :: complex | float, b :: complex | float) :: complex
  def sub({r_a, i_a}, {r_b, i_b}) do
    {r_a - r_b, i_a - i_b}
  end
  def sub({r_a, i_a}, b), do: {r_a - b, i_a}
  def sub(a, {r_b, i_b}), do: {a - r_b, i_b * -1}
  

  @doc """
  Divide two complex numbers, or a real and a complex number
  """
  @spec div(a :: complex | float, b :: complex | float) :: complex
  def div({r_a, i_a} = a, {r_b, i_b} = b) do
    {(real(a) * real(b) + imaginary(a) * imaginary(b)) / (real(b) * real(b) + imaginary(b) * imaginary(b)) , (imaginary(a) * real(b) - real(a) * imaginary(b)) / (real(b) * real(b) + imaginary(b) * imaginary(b))}
  end

  def div({r_a, i_a}, b), do: {r_a / b, i_a / b}
  def div(a, {r_b, i_b}), do: { (a*r_b + i_b * 0)/(r_b*r_b + i_b*i_b), (0 * r_b - a * i_b)/(r_b*r_b + i_b*i_b)  }

  @doc """
  Absolute value of a complex number
  """
  @spec abs(a :: complex) :: float
  def abs(a) do
    :math.sqrt(real(a) * real(a) + imaginary(a) * imaginary(a))
  end

  @doc """
  Conjugate of a complex number
  """
  @spec conjugate(a :: complex) :: complex
  def conjugate(a) do
    {real(a), imaginary(a) * -1}
  end

  @doc """
  Exponential of a complex number
  """
  @spec exp(a :: complex) :: complex
  def exp(a) do
    {:math.exp(real(a)) * :math.cos(imaginary(a)), :math.exp(real(a)) * :math.sin(imaginary(a))}
  end
end
