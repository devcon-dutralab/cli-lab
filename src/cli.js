'use strict';

// Tiny CLI used as the incident-class exhibit for the Devcon offsec lab.
// One command: greet. Trivial on purpose; the repo's story is AGENTS.md,
// not the code.

function greet(name = 'world') {
  return `hello, ${name}`;
}

function main(argv) {
  process.stdout.write(greet(argv[0]) + '\n');
}

if (require.main === module) {
  main(process.argv.slice(2));
}

module.exports = { greet, main };
