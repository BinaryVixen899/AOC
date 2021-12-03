# require 'IO'
text = File.open('data.txt').read
increasecount = 0
previousvalue = 0

File.foreach('data.txt').with_index do |line, line_no|
  currentvalue = Integer(line)
  if line_no != 0 
    if currentvalue > previousvalue
      increasecount += 1    
      previousvalue = currentvalue
    else
      previousvalue = currentvalue
    end

  else
    previousvalue = currentvalue  
  end
  
  puts "increase count: #{increasecount}"
end
