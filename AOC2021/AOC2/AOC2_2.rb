require 'pry' 
$position = 0 
$depth = 0 
$aim = 0 

course = [
"forward 5", 
"down 5", 
"forward 8",
"up 3", 
"down 8",
"forward 2"]
course = File.readlines("data.txt")

# course = course.split(' ').strip

def Calculate(direction, amount)
  amount = amount.to_i
  # if ['f', 'b', 'u', 'd'].include?(direction)
  #   direction = direction
  # end
  puts amount


  case direction

  when 'f'
    # binding.pry
    $position += amount
    $depth += (amount * $aim)
  
  when 'u' 
    $aim += (amount * -1)

  when 'd'
    $aim += amount
 
  else
    puts "Oh so there is a backwards"
    
  end
end



for i in course
  direction = i[0]
  i = i.split(" ")
  amount = i[1]
  Calculate(direction, amount)
end

puts ($position * $depth)