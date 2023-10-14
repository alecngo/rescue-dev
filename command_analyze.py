import pprint 
import gpt4all
 
model = gpt4all.GPT4All("orca-mini-7b.ggmlv3.q4_0.bin")
with model.chat_session():
    response = model.generate("Where is Berea College")
    print(response)
    pprint.pprint(model.current_chat_session)