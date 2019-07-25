# Yin
- Bytes are stored in the red channel (using RGBA colourscheme) of even "x-coordinate" pixles starting from 0
    - i.e. Pixels (0, 0), (2, 0), (4, 0), ...

- Bytes will spell out the flag in ASCII

- I am indeed a Ying main in Rainbow 6 Siege just in case you were wondering

# Yang
- Bytes are stored in the least significant digits of the RGB channels in decimal of every odd coordinate
    - i.e. Pixels (1, 1), (1, 3), (1, 5), ...

- Example:
    - let us encode the letter 'A' in the pixel of RGBA value `[100, 50, 0, 255]`
    - encoding letter 'A' is 0x41 in hex (as you all know already) and 65 in decimal.
    - this translates to 065 so the encoded pixel looks like `[100, 54, 1, 255]`
    - decoding is just reversing this

# YinYang
- The message is encoded into the image using the same scheme as Yang, but starting at (444, 444)
    - The encoding I used adds on top of the original image 
    - It uses the same encoding scheme as Yang, but for every even pixel, just like Yin
        - Hence YinYang...

- However, original image is not technically required, it is a red herring in this case
    - Message is encoded into the black portions of the image, making it the same as Yang
    - If you use the script to encode a message, you will get a much harder steg than this challenge!

- The original image's URL is encoded using Yin's encoding scheme and is malformed, another red herring
    - Original image is a .svg
    - Needs to be converted to a png of the correct size 
    - Use your favourite image manipulation program (I used GIMP)

# Concluding points
- I chose to challenges MUCH easier than O-Week's steg, given the time constraints of this CTF
    - In a 24-hour or more CTF, I will include much more challenging stegs >:D

- Let me know if you liked this style of progressive challenges! I'm curious to know what you guys liked and didn't like about my challenges.

- I really enjoyed creating these challenges, I hope those of you who solved it enjoyed solving it :)

- Feedback is appreciated, please direct those for this and other challenges to SecSoc, we will route them to the authors.
