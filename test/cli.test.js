'use strict';

const test = require('node:test');
const assert = require('node:assert');
const { greet } = require('../src/cli');

test('greet names the target', () => {
  assert.equal(greet('devcon'), 'hello, devcon');
});

test('greet defaults to world', () => {
  assert.equal(greet(), 'hello, world');
});
