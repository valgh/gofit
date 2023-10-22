# gofit

**gofit** is a CLI Utility built with Golang which invokes ChatGPT's endpoint in order to send user-defined requests. **gofit** is at its **rel-0.1** versions, which is still a draft and work in progress: many changes are still to be implemented, such as better command handling, new features and unit tests/exception handling.

**gofit**'s goal is that of helping the user requesting a workout plan to ChatGPT, by either leveraging a YAML config file for its workout, or by providing as input the user and training details.

**WARNING**: **gofit** does not bill the user for its usage, but in order to request ChatGPT a workout plan, an APIKEY must be provided through 'gofit --cfggpt' command. The user musy have a valid ChatGPT account with billing details, and billing will be based upon the plan the user chooses on its own ChatGPT account. The **max number of tokens** parameter can be set in the ChatGPT config file to truncate ChatGPT reponses and limit the billing, as well as ChatGPT billing details, thresholds and limits which can be set in the account.

See [here](https://platform.openai.com/docs/guides/gpt) for more information.

**WARNING**: gofit does not retain user data, but all the data provided in config files or input to the 'workout' command will be used by ChatGPT according to its settings. 
See [here](https://help.openai.com/en/articles/7730893-data-controls-faq) for more information.

# Usage

# What's next
 
