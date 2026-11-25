// ESLint flat config for the JavaScript examples.
export default [
  {
    rules: {
      "no-unused-vars": "error",
      "no-restricted-syntax": [
        "error",
        {
          selector: "Identifier[name='secretSauce']",
          message: "CR-MARKER-T1-RULE: restricted identifier used",
        },
      ],
    },
  },
];
