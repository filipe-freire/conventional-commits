use inquire::validator::Validation;
use inquire::{Select, Text};
use std::collections::HashMap;
use std::process::{Command, exit};

fn get_commit_values() -> HashMap<&'static str, &'static str> {
    let mut values = HashMap::new();
    values.insert("feat", "💎 feat:");
    values.insert("add", "🎁 add:");
    values.insert("update", "🆙 update:");
    values.insert("ref", "🔧 ref:");
    values.insert("wip", "⏳ wip:");
    values.insert("delete", "🔥 delete:");
    values.insert("chore", "🧹 chore:");
    values.insert("bugfix", "🐛 bugfix:");
    values
}

// Define commit options for display
struct CommitOption {
    key: &'static str,
    label: &'static str,
}

impl std::fmt::Display for CommitOption {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.label)
    }
}

fn get_commit_options() -> Vec<CommitOption> {
    vec![
        CommitOption {
            key: "feat",
            label: "💎 Feature",
        },
        CommitOption {
            key: "add",
            label: "🎁 Minor improvement",
        },
        CommitOption {
            key: "update",
            label: "🆙 Update",
        },
        CommitOption {
            key: "ref",
            label: "🔧 Refactor",
        },
        CommitOption {
            key: "wip",
            label: "⏳ Work In Progress",
        },
        CommitOption {
            key: "delete",
            label: "🔥 Deletion",
        },
        CommitOption {
            key: "chore",
            label: "🧹 Chore",
        },
        CommitOption {
            key: "bugfix",
            label: "🐛 Bugfix",
        },
    ]
}

fn main() {
    println!("Select your commit type:");

    let commit_values = get_commit_values();
    let commit_options = get_commit_options();

    // Display selection prompt using inquire
    let selection = Select::new("Select your commit type:", commit_options)
        .with_help_message("Use arrow keys to navigate, Enter to select")
        .prompt();

    let selected_option = match selection {
        Ok(option) => option,
        Err(_) => {
            println!("Error selecting commit type");
            exit(1);
        }
    };

    // Get the selected commit type
    let commit_type = commit_values.get(selected_option.key).unwrap();

    // Prompt for commit message using inquire
    let msg = Text::new("Commit message:")
        .with_help_message("write a clear and concise message that describes the changes made")
        .with_validator(|value: &str| {
            if value.trim().is_empty() {
                Err("Please enter a commit message".into())
            } else {
                Ok(Validation::Valid)
            }
        })
        .prompt();

    let msg = match msg {
        Ok(message) => message,
        Err(_) => {
            println!("Error getting commit message");
            exit(1);
        }
    };

    // Execute git add .
    let add_status = Command::new("git").args(&["add", "."]).status();

    if let Err(e) = add_status {
        println!("Failed to execute git add command: {}", e);
        exit(1);
    }

    let add_status = add_status.unwrap();
    if !add_status.success() {
        println!("Something went wrong with git add! 😿");
        exit(1);
    }

    // Execute git commit
    let commit_message = format!("{} {}", commit_type, msg);
    let commit_status = Command::new("git")
        .args(&["commit", "-m", &commit_message])
        .status();

    if let Err(e) = commit_status {
        println!("Failed to execute git commit command: {}", e);
        exit(1);
    }

    let commit_status = commit_status.unwrap();
    if commit_status.success() {
        println!("Done! 💪");
        exit(0);
    } else {
        println!("Something went wrong! 😿");
        exit(1);
    }
}
