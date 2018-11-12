import std from 'std';

// We can print basic types to stdout.
std.log(1.2);
std.log('foo');
// XXX std.log(undefined);
std.log(null);

// We can print an object to stdout.
std.log({ kind: 'Bar', foo: 1.2 });

// Key order is deterministic.
std.log({ foo: 1.2, kind: 'Bar' });

// Same, but in YAML
std.log({ kind: 'Bar', foo: { number: 1.2, string: 'mystring' } }, std.Format.YAML);
std.log({ foo: { number: 1.2, string: 'mystring' }, kind: 'Bar' }, std.Format.YAML);
