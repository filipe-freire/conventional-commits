import { intro, outro, select, text } from "npm:@clack/prompts";
import { COMMIT_OPTIONS, COMMIT_VALUES } from "./constants.ts";

async function main() {
	intro("Select your commit type:");

	const commitType = await select({
		message: "Select your commit type:",
		initialValue: COMMIT_VALUES.feat,
		options: COMMIT_OPTIONS,
	});

	const msg = await text({
		message: "Commit message:",
		placeholder: "short and sweet 🍩",
		validate: (value) => {
			if (!value.trim()) {
				return "Please enter a commit message";
			}
		},
	});

	new Deno.Command("git", {
		args: ["add", "."],
		stdin: "piped",
		stdout: "piped",
	}).spawn();

	// Execute command in bash through Deno
	const cmd = new Deno.Command("git", {
		args: ["commit", "-m", `${commitType.toString()} ${msg.toString()}`],
		stdin: "piped",
		stdout: "piped",
	});

	const status = await cmd.spawn().status;

	if (status.success) {
		outro("Done! 💪");
		return Deno.exit(0);
	}

	outro("Something went wrong! 😿");

	Deno.exit(1);
}

main();
