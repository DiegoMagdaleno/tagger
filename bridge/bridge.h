//
//  bridge.m
//  tagger
//
//  Created by Diego Magdaleno on 10/30/20.
//

#ifndef bridge_h
#define bridge_h

#import <Foundation/Foundation.h>

const char* NSStringToCString(NSString*);
const NSArray* getFilesWithCertainMacOSTag(char const*);
const void* NSArrayItem(NSArray*, unsigned long);
unsigned long NSArrayLen(NSArray*);
int NSNumberToInt(NSNumber*);

#endif /* bridge_h */