import * as std from '@jkcfg/std';

const value = std.read('./test-jsonstream.js.expected', { format: std.Format.JSONStream });
value.then(v => std.write(v, '', { format: std.Format.JSONStream }));
