# An-Exercise-in-Automated-Programming
Introduction: 
The purpose of this documentation is to provide a comprehensive overview of the evaluation of automated programming tools for a technology startup. As an outside consultant, the goal was to assess the potential of automated code generation, AI-assisted programming, and AI-generated code in reducing programming workload and costs.

Methodology:
The evaluation process involved the following steps:
•	Selection of Tools:
•	Identified code generators in the Go ecosystem (e.g., standard library generate, jennifer package).
•	Chose GitHub Copilot for AI-assisted programming but since Github Copilot is not free, I used Codeium. Codeium is the modern coding superpower, a free code acceleration toolkit built on cutting edge AI technology. Currently, Codeium provides code completion tool in over 70+ languages, with lightning fast speeds and state-of-the-art suggestion quality.
•	Selected ChatGPT as the LLM-based service for AI-generated code. Provided it with a prompt and it generated code how I wanted.
•	Application to Assignment: 
•	Applied each method to the ‘Go For Statistics’ assignment, which involved statistical analysis using the Anscombe Quartet data.
Automated Code Generation:
Tools Examined:
•	Go standard library generate
•	jennifer package

Approach:
•	Explored the capabilities of code generators to automate repetitive tasks.
•	Investigated the possibility of generating unit tests for statistical functions.

Code Samples:
•	Example of code snippet generated for unit test: Automated_Code_Generation/test_generator.go and Automated_Code_Generation/main_test.go.
•	Test_generator.go has the logic defined for generating test units for the main.go using Jennifer package. In this code, I used Jennifer.

AI-assisted Programming:
Tool Used:
•	Codeium
Approach:
•	Integrated Codeium with VS Code for real-time assistance.
•	Utilized Codeium’s suggestions to revise existing code and improve efficiency.
Results:
•	Improved readability and efficiency of code snippets through Codeium’s recommendations.
•	Example of before-and-after code snippets: AI_Assisted_Programming/main.go is before and AI_Assisted_Programming/main_revised.go is after.

AI-generated Code:
Service Used:
•	ChatGPT (on free plan)
Approach:
•	Engaged in a dialogue with ChatGPT to generate Go code for statistical analysis.
•	Explored various prompts and evaluated the quality of generated code.
Results:
•	Generated code snippets for statistical functions based on ChatGPT's responses.
•	Example conversation with ChatGPT: Generate Go code for linear regression analysis on the Anscombe data set. Make it clean, optimized and readable.
Evaluation:
Effectiveness:
•	Automated code generation tools provided some level of assistance in streamlining development tasks.
•	Codeium significantly improved code quality and productivity by providing context-aware suggestions.
•	ChatGPT demonstrated the potential to generate code snippets but may require refinement for specific use cases.
Limitations:
•	Code generators may not cover all edge cases and require manual intervention for complex scenarios.
•	AI-assisted programming tools like Copilot and Codeium may lack domain-specific knowledge and provide generic suggestions.
•	AI-generated code from ChatGPT may not always meet the requirements and require human validation.
Recommendations:
Based on the evaluation, the following recommendations are proposed:
•	Implement GitHub Copilot and Codeium etc for AI-assisted programming to enhance productivity and code quality.
•	Use automated code generation tools selectively for repetitive tasks but maintain human oversight for complex scenarios.
•	Explore further refinement of AI-generated code tools to align with specific project requirements.
Conclusion:
The evaluation of automated programming tools revealed promising opportunities to enhance the efficiency of software development processes. By leveraging AI-assisted programming and code generation tools, the startup can optimize resource utilization and accelerate product development while maintaining code quality.
References:
•	GitHub Copilot  
•	ChatGPT  
•	Go standard library 
•	Codeium
•	Jennifer package for Go
