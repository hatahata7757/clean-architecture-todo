module.exports = {
  extends: ['@commitlint/config-conventional'],
  rules: {
    'type-enum': [
      2,
      'always',
      [
        'add',
        'fix',
        'refactor',
        'revert',
        'update',
        'remove'
      ]
    ]
  }
};
