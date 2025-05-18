export const COMMIT_VALUES = {
	feat: "💎 feat:",
	add: "🎁 add:",
	update: "🆙 update:",
	ref: "🔧 ref:",
	wip: "⏳ wip:",
	delete: "🔥 delete:",
	chore: "🧹 chore:",
	bugfix: "🐛 bugfix:",
} as const;

export const COMMIT_OPTIONS = [
	{ value: COMMIT_VALUES.feat, label: "💎 Feature" },
	{ value: COMMIT_VALUES.add, label: "🎁 Minor improvement" },
	{ value: COMMIT_VALUES.update, label: "🆙 Update" },
	{ value: COMMIT_VALUES.ref, label: "🔧 Refactor" },
	{ value: COMMIT_VALUES.wip, label: "⏳ Work In Progress" },
	{ value: COMMIT_VALUES.delete, label: "🔥 Deletion" },
	{ value: COMMIT_VALUES.chore, label: "🧹 Chore" },
	{ value: COMMIT_VALUES.bugfix, label: "🐛 Bugfix" },
];
