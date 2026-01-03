<?php

/*
 * By adding type hints and enabling strict type checking, code can become
 * easier to read, self-documenting and reduce the number of potential bugs.
 * By default, type declarations are non-strict, which means they will attempt
 * to change the original type to match the type specified by the
 * type-declaration.
 *
 * In other words, if you pass a string to a function requiring a float,
 * it will attempt to convert the string value to a float.
 *
 * To enable strict mode, a single declare directive must be placed at the top
 * of the file.
 * This means that the strictness of typing is configured on a per-file basis.
 * This directive not only affects the type declarations of parameters, but also
 * a function's return type.
 *
 * For more info review the Concept on strict type checking in the PHP track
 * <link>.
 *
 * To disable strict typing, comment out the directive below.
 */

declare(strict_types=1);

class Robot
{
    static $nameInUse = [];
    private $name;

    public function __construct()
    {
        $this->name = $this->generateName();
    }

    public function getName(): string
    {
        return $this->name;
    }

    public function reset(): void
    {
        $this->name = $this->generateName();
    }

    public function generateRandomCharacters($length = 2) : string
    {
        $leters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
        $letersLength = strlen($leters);
        $randomString = '';
        for ($i = 0; $i < $length; $i++) {
            $randomString .= $leters[rand(0, $letersLength - 1)];
        }
        return $randomString;
    }

    public function generateRandomNumbers($length = 3) : string
    {
        $numbers = '0123456789';
        $numbersLength = strlen($numbers);
        $randomString = '';
        for ($i = 0; $i < $length; $i++) {
            $randomString .= $numbers[rand(0, $numbersLength - 1)];
        }
        return $randomString;
    }

    private function generateName()
    {
        do {
            $name = $this->generateRandomCharacters(2) . $this->generateRandomNumbers(3);
        } while(in_array($name, self::$nameInUse));
        array_push(self::$nameInUse, $name);
        return $name;
    }
}
