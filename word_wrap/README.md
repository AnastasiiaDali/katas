Create a function to "wrap" a string around a certain line width, similar to how a text editor does it.

Your input will be a number of paragraphs, each on one long line, separated from each other by a blank line; as well as a max line width.

Your output should be a number of lines, each of which should not exceed the max line width.

Sample input:

```The prize was delivered to Tom with as much effusion as the superintendent could pump up under under the circumstances; but it lacked somewhat of the true gush, for the poor fellow’s instinct taught him that there was a mystery here that could not well bear the light, perhaps; it was simply preposterous that this boy had warehoused two thousand sheaves of Scriptural wisdom on his premises—a dozen would strain his capacity, without a doubt. \n\n Amy Lawrence was proud and glad, and she tried to make Tom see it in her face—but he wouldn’t look. She wondered; then she was just a grain troubled; next a dim suspicion came and went—came again; she watched; a furtive glance told her worlds—and then her heart broke, and she was jealous, and angry, and the tears came and she hated everybody. Tom most of all (she thought).```


Sample output (line width 58 chars):

```The prize was delivered to Tom with as much effusion as
the superintendent could pump up under the circumstances;
but it lacked somewhat of the true gush, for the poor
fellow’s instinct taught him that there was a mystery here
that could not well bear the light, perhaps; it was simply
preposterous that this boy had warehoused two thousand
sheaves of Scriptural wisdom on his premises—a dozen
would strain his capacity, without a doubt.

Amy Lawrence was proud and glad, and she tried to make
Tom see it in her face—but he wouldn’t look. She wondered;
then she was just a grain troubled; next a dim suspicion
came and went—came again; she watched; a furtive glance
told her worlds—and then her heart broke, and she was
jealous, and angry, and the tears came and she hated
everybody. Tom most of all (she thought).```
